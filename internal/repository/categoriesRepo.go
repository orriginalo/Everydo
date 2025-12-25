package repository

import (
	"Everydo/internal/models"

	"gorm.io/gorm"
)

type ICategoriesRepository interface {
	CreateCategory(category *models.Category) error
	GetCategories() []models.Category
	DeleteCategory(id int)
}

type CategoriesRepository struct {
	db *gorm.DB
}

func NewCategoriesRepository(db *gorm.DB) ICategoriesRepository {
	return &CategoriesRepository{
		db: db,
	}
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
