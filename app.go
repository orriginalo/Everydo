package main

import (
	"Everydo/internal/db"
	"Everydo/internal/models"
	"Everydo/internal/repository"
	"Everydo/internal/utils"
	"context"
	"log/slog"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Repositories struct {
	tasksRepo      repository.ITasksRepository
	categoriesRepo repository.ICategoriesRepository
}

type App struct {
	repo Repositories
	ctx  context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {

	db := db.InitDB("data.db")
	// db := db.InitDB(utils.GetDataPath())
	repositories := Repositories{
		tasksRepo:      repository.NewTasksRepository(db),
		categoriesRepo: repository.NewCategoriesRepository(db),
	}

	tasks := repositories.tasksRepo.GetAllTasks()
	for _, task := range tasks {
		newTime, valid := utils.CheckNextResetValid(task)
		if !valid {
			repositories.tasksRepo.UpdateTask(int(task.ID), map[string]interface{}{
				"next_reset_at": newTime,
			})
		}
		slog.Info("Все задания проверены")
	}
	a.repo = repositories
	a.ctx = ctx
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

	err := a.repo.categoriesRepo.UpdateCategory(id, updates)
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

	err := a.repo.categoriesRepo.CreateCategory(category)
	if err != nil {
		return 0, err
	}

	return category.ID, nil
}

func (a *App) GetCategories() []models.Category {
	categories := a.repo.categoriesRepo.GetCategories()
	return categories
}

func (a *App) DeleteCategory(id int) {
	a.repo.categoriesRepo.DeleteCategory(id)
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

	err := a.repo.tasksRepo.CreateTask(&task)
	if err != nil {
		return 0, err
	}

	return task.ID, nil
}

func (a *App) GetTasks(categoryID int) []models.Task {
	tasks := a.repo.tasksRepo.GetTasks(categoryID)
	return tasks
}

func (a *App) CompleteTask(id int) {
	task := a.repo.tasksRepo.GetTask(id)

	now := time.Now()
	nextReset := utils.CalcNextReset(task, now)

	a.repo.tasksRepo.UpdateTask(id, map[string]interface{}{
		"is_completed":  true,
		"last_done_at":  now,
		"next_reset_at": nextReset,
	})
}

func (a *App) UncompleteTask(id int) {
	a.repo.tasksRepo.UpdateTask(id, map[string]interface{}{
		"is_completed": false,
	})
}

func (a *App) UpdateNextReset(id int) {
	task := a.repo.tasksRepo.GetTask(id)
	nextReset := utils.CalcNextReset(task, time.Now())
	a.repo.tasksRepo.UpdateTask(id, map[string]interface{}{
		"next_reset_at": nextReset,
	})
}

func (a *App) DeleteTask(id int) {
	a.repo.tasksRepo.DeleteTask(id)
}

func (a *App) UpdateTask(
	id int,
	name string,
	reloadType string,
	reloadEvery int,
	resetTime string,
	resetWeekday *int,
) {
	a.repo.tasksRepo.UpdateTask(id, map[string]interface{}{
		"name":          name,
		"reload_type":   reloadType,
		"reload_every":  reloadEvery,
		"reset_time":    resetTime,
		"reset_weekday": resetWeekday,
	})
	a.repo.tasksRepo.UpdateTask(id, map[string]interface{}{
		"next_reset_at": utils.CalcNextReset(a.repo.tasksRepo.GetTask(id), time.Now()),
	})
}
