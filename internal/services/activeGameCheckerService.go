package services

import (
	"Everydo/internal/errs"
	"Everydo/internal/models"
	"Everydo/internal/repository"
	"slices"

	"github.com/mitchellh/go-ps"
)

type GameChecker struct {
	repo *repository.Repositories
}

func NewGameChecker(repo *repository.Repositories) *GameChecker {
	return &GameChecker{repo: repo}
}

func (g *GameChecker) GetActiveCategories() ([]models.Category, error) {
	processesExeNames := []string{}

	processes, err := ps.Processes()
	if err != nil {
		return []models.Category{}, err
	}
	for _, p := range processes {
		processesExeNames = append(processesExeNames, p.Executable())
	}

	categories := g.repo.CategoriesRepo.GetCategories()

	activeCategories := []models.Category{}
	for _, c := range categories {
		if slices.Contains(processesExeNames, c.ExeName) {
			activeCategories = append(activeCategories, c)
		}
	}
	if len(activeCategories) == 0 {
		return []models.Category{}, errs.ErrNoActiveCategory
	}
	return activeCategories, nil
}
