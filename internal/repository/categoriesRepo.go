package repository

import (
	"Everydo/internal/models"

	"gorm.io/gorm"
)

type ICategoriesRepository interface {
	CreateCategory(category *models.Category) error
	GetCategories() []models.Category
	GetCategoryByExeName(exeName string) *models.Category
	DeleteCategory(id int)
	UpdateCategory(id int, updates map[string]interface{}) error
}

type CategoriesRepository struct {
	db *gorm.DB
}

func NewCategoriesRepository(db *gorm.DB) ICategoriesRepository {
	return &CategoriesRepository{
		db: db,
	}
}

func (r *CategoriesRepository) GetCategoryByExeName(exeName string) *models.Category {
	var category models.Category
	r.db.First(&category, "exe_name = ?", exeName)
	return &category
}

func (r *CategoriesRepository) UpdateCategory(id int, updates map[string]interface{}) error {
	return r.db.Model(&models.Category{}).Where("id = ?", id).Updates(updates).Error
}

func (r *CategoriesRepository) CreateCategory(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoriesRepository) GetCategories() []models.Category {
	var categories []models.Category
	r.db.Find(&categories)
	return categories
}

func (r *CategoriesRepository) DeleteCategory(id int) {
	r.db.Delete(&models.Category{}, id)
}
