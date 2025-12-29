package main

import (
	"Everydo/internal/db"
	"Everydo/internal/errs"
	"Everydo/internal/models"
	"Everydo/internal/repository"
	"Everydo/internal/services"
	"Everydo/internal/utils"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"slices"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	repo repository.Repositories
	ctx  context.Context

	activeCategoriesTimes map[string]int
}

func NewApp() *App {
	return &App{
		activeCategoriesTimes: make(map[string]int),
	}
}

func (a *App) startup(ctx context.Context) {

	db := db.InitDB("data.db")
	// db := db.InitDB(utils.GetDataPath())
	repositories := repository.Repositories{
		TasksRepo:      repository.NewTasksRepository(db),
		CategoriesRepo: repository.NewCategoriesRepository(db),
	}
	notifier := services.NewNotifier(&repositories)
	activeGameChecker := services.NewGameChecker(&repositories)

	ticker := services.NewTicker(
		notifier, time.Second,
		func() {
			tickerFunc(
				a, activeGameChecker,
				nil, func(c *models.Category) {},
			)
		})
	ticker.Start()

	tasks := repositories.TasksRepo.GetAllTasks()
	for _, task := range tasks {
		newTime, valid := utils.CheckNextResetValid(task)
		if !valid {
			repositories.TasksRepo.UpdateTask(int(task.ID), map[string]interface{}{
				"next_reset_at": newTime,
				"is_completed":  false,
			})
		}
	}
	slog.Info("Все задания проверены")
	a.repo = repositories
	a.ctx = ctx
}

func tickerFunc(
	app *App, activeGameChecker *services.GameChecker,
	onEnterCategory func(*models.Category),
	onExitCategory func(*models.Category),
) {
	slog.Info("Tick --------------------")
	fmt.Println(app.activeCategoriesTimes)

	activeCategories, err := activeGameChecker.GetActiveCategories()
	if err != nil {
		if errors.Is(err, errs.ErrNoActiveCategory) {
			for k := range app.activeCategoriesTimes {
				delete(app.activeCategoriesTimes, k)
			}
		} else {
			slog.Error(err.Error())
		}
		return
	}

	activeCategoriesExeNames := []string{}

	for _, c := range activeCategories {
		activeCategoriesExeNames = append(activeCategoriesExeNames, c.ExeName)
		if _, ok := app.activeCategoriesTimes[c.ExeName]; !ok {
			app.activeCategoriesTimes[c.ExeName] = 0
			onEnterCategory(&c)
		}
		app.activeCategoriesTimes[c.ExeName]++
	}

	for k := range app.activeCategoriesTimes {
		if !slices.Contains(activeCategoriesExeNames, k) {
			delete(app.activeCategoriesTimes, k)

			category := app.repo.CategoriesRepo.GetCategoryByExeName(k)
			onExitCategory(category)
		}
	}

}

// Categories

func (a *App) OpenURL(url string) {
	runtime.BrowserOpenURL(a.ctx, url)
}

func (a *App) UpdateCategory(id int, name, exeName string) error {
	updates := map[string]interface{}{
		"name":     name,
		"exe_name": exeName,
	}

	err := a.repo.CategoriesRepo.UpdateCategory(id, updates)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) CreateCategory(name, exeName string) (uint, error) {
	category := &models.Category{
		Name:    name,
		ExeName: exeName,
	}

	err := a.repo.CategoriesRepo.CreateCategory(category)
	if err != nil {
		return 0, err
	}

	return category.ID, nil
}

func (a *App) GetCategories() []models.Category {
	categories := a.repo.CategoriesRepo.GetCategories()
	return categories
}

func (a *App) DeleteCategory(id int) {
	a.repo.CategoriesRepo.DeleteCategory(id)
}

// Tasks

func (a *App) CreateTask(
	categoryID int,
	name string,
	reloadType string,
	reloadEvery int,
	resetTime string,
	resetWeekday *int,
) (uint, error) {

	task := models.Task{
		CategoryID:   uint(categoryID),
		Name:         name,
		ReloadType:   reloadType,
		ReloadEvery:  reloadEvery,
		ResetTime:    resetTime,
		ResetWeekday: resetWeekday,
	}

	task.NextResetAt = utils.CalcNextReset(task, time.Now())

	err := a.repo.TasksRepo.CreateTask(&task)
	if err != nil {
		return 0, err
	}

	return task.ID, nil
}

func (a *App) GetTasks(categoryID int) []models.Task {
	tasks := a.repo.TasksRepo.GetTasks(categoryID)
	return tasks
}

func (a *App) CompleteTask(id int) {
	task := a.repo.TasksRepo.GetTask(id)

	now := time.Now()
	nextReset := utils.CalcNextReset(task, now)

	a.repo.TasksRepo.UpdateTask(id, map[string]interface{}{
		"is_completed":  true,
		"last_done_at":  now,
		"next_reset_at": nextReset,
	})
}

func (a *App) UncompleteTask(id int) {
	a.repo.TasksRepo.UpdateTask(id, map[string]interface{}{
		"is_completed": false,
	})
}

func (a *App) UpdateNextReset(id int) {
	task := a.repo.TasksRepo.GetTask(id)
	nextReset := utils.CalcNextReset(task, time.Now())
	a.repo.TasksRepo.UpdateTask(id, map[string]interface{}{
		"next_reset_at": nextReset,
	})
}

func (a *App) DeleteTask(id int) {
	a.repo.TasksRepo.DeleteTask(id)
}

func (a *App) UpdateTask(
	id int,
	name string,
	reloadType string,
	reloadEvery int,
	resetTime string,
	resetWeekday *int,
) {
	a.repo.TasksRepo.UpdateTask(id, map[string]interface{}{
		"name":          name,
		"reload_type":   reloadType,
		"reload_every":  reloadEvery,
		"reset_time":    resetTime,
		"reset_weekday": resetWeekday,
	})
	a.repo.TasksRepo.UpdateTask(id, map[string]interface{}{
		"next_reset_at": utils.CalcNextReset(a.repo.TasksRepo.GetTask(id), time.Now()),
	})
}
