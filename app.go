package main

import (
	"Everydo/internal/db"
	"Everydo/internal/models"
	"Everydo/internal/repository"
	"Everydo/internal/utils"
	"context"
	"time"
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
	repositories := Repositories{
		tasksRepo:      repository.NewTasksRepository(db),
		categoriesRepo: repository.NewCategoriesRepository(db),
	}

	a.repo = repositories
	a.ctx = ctx
}

// Categories

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

	err := a.repo.tasksRepo.CreateTask(&task)
	if err != nil {
		return 0, err
	}

	return task.ID, nil
}

func (a *App) GetTasks(categoryID int) map[string][]models.Task {
	tasksByType := map[string][]models.Task{}

	dailyTasks := a.repo.tasksRepo.GetTasksWithType(categoryID, "daily")
	weeklyTasks := a.repo.tasksRepo.GetTasksWithType(categoryID, "weekly")
	customTasks := a.repo.tasksRepo.GetTasksWithType(categoryID, "custom")

	if len(dailyTasks) > 0 {
		tasksByType["daily"] = dailyTasks
	}
	if len(weeklyTasks) > 0 {
		tasksByType["weekly"] = weeklyTasks
	}
	if len(customTasks) > 0 {
		tasksByType["custom"] = customTasks
	}

	return tasksByType
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
