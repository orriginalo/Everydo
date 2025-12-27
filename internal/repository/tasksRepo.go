package repository

import (
	"Everydo/internal/models"

	"gorm.io/gorm"
)

type ITasksRepository interface {
	CreateTask(task *models.Task) error
	GetAllTasks() []models.Task
	GetTasks(categoryID int) []models.Task
	GetTasksWithType(categoryID int, reloadType string) []models.Task
	GetTask(id int) models.Task
	UpdateTask(id int, updates map[string]interface{})
	DeleteTask(id int)
}

type TasksRepository struct {
	db *gorm.DB
}

func NewTasksRepository(db *gorm.DB) ITasksRepository {
	return &TasksRepository{
		db: db,
	}
}

func (r *TasksRepository) GetAllTasks() []models.Task {
	var tasks []models.Task
	r.db.Find(&tasks)
	return tasks
}

func (r *TasksRepository) CreateTask(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *TasksRepository) GetTasksWithType(categoryID int, reloadType string) []models.Task {
	var tasks []models.Task
	r.db.Model(&models.Task{}).Where("category_id = ? AND reload_type = ?", categoryID, reloadType).Find(&tasks)
	return tasks
}

func (r *TasksRepository) GetTasks(categoryID int) []models.Task {
	var tasks []models.Task
	r.db.Model(&models.Task{}).Where("category_id = ?", categoryID).Find(&tasks)
	return tasks
}

func (r *TasksRepository) GetTask(id int) models.Task {
	var task models.Task
	r.db.First(&task, id)
	return task
}

func (r *TasksRepository) UpdateTask(id int, updates map[string]interface{}) {
	r.db.Model(&models.Task{}).Where("id = ?", id).Updates(updates)
}

func (r *TasksRepository) DeleteTask(id int) {
	r.db.Delete(&models.Task{}, id)
}
