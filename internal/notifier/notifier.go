package notifier

import (
	"Everydo/internal/models"
	"Everydo/internal/repository"
	_ "embed"
	"fmt"
	"log/slog"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/mitchellh/go-ps"
)

//go:embed ..\..\build\appicon.png
var iconData []byte

type Notifier struct {
	repo                  *repository.Repositories
	activeCategoryProcess *string
}

func NewNotifier(repo *repository.Repositories) *Notifier {
	return &Notifier{
		repo:                  repo,
		activeCategoryProcess: nil,
	}
}

func (n *Notifier) Start() {
	ticker := time.NewTicker(5 * time.Second)
	slog.Info("Ticker started")
	go func() {
		for range ticker.C {
			n.Check()
		}
	}()
}

func (n *Notifier) Stop() {

}

func (n *Notifier) Check() {
	processes, err := ps.Processes()
	if err != nil {
		panic(err)
	}

	categories := n.repo.CategoriesRepo.GetCategories()

	isProcessActive := false
	for _, c := range categories {
		for _, p := range processes {
			if p.Executable() == c.ExeName {
				isProcessActive = true
				if n.activeCategoryProcess != nil && *n.activeCategoryProcess == c.ExeName {
					continue
				}

				n.activeCategoryProcess = &c.ExeName
				n.Notify(c)
				return
			}
		}
	}
	if n.activeCategoryProcess != nil && !isProcessActive {
		n.activeCategoryProcess = nil
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
