package services

import (
	"Everydo/internal/models"
	"Everydo/internal/repository"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/gen2brain/beeep"
)

//go:embed ..\..\build\appicon.png
var iconData []byte

type Notifier struct {
	repo                  *repository.Repositories
	activeCategoryProcess *string
	// activeCategorySeconds int
}

func NewNotifier(repo *repository.Repositories) *Notifier {
	return &Notifier{
		repo:                  repo,
		activeCategoryProcess: nil,
	}
}

func (n *Notifier) Notify(category models.Category) {
	tasks := n.repo.TasksRepo.GetTasks(int(category.ID))

	var uncompletedTasksCount int
	for _, task := range tasks {
		if !task.IsCompleted {
			uncompletedTasksCount++
		}
	}

	if uncompletedTasksCount == 0 {
		return
	}

	title := fmt.Sprintf("Остались невыполненные задачи в %v", category.Name)
	msg := fmt.Sprintf("Ежедневные: %v", uncompletedTasksCount)

	if err := beeep.Notify(title, msg, iconData); err != nil {
		slog.Error(err.Error())
	}
}
