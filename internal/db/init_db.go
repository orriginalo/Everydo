package db

import (
	"Everydo/internal/models"
	"log/slog"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitDB(dbFilename string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbFilename), &gorm.Config{})
	if err != nil {
		slog.Error("Не удалось подключиться к БД: " + err.Error())
	}

	slog.Info("БД: Подключение | Успешно.")
	if err := db.AutoMigrate(&models.Category{}, &models.Task{}); err != nil {
		slog.Error("Не удалось произвести миграцию: " + err.Error())
	}
	slog.Info("БД: Автомиграция | Успешно.")
	return db
}
