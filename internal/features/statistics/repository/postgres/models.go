package statistics_postgres_repository

import (
	"time"

	"github.com/mishalyamets1/golang_todoapp/internal/core/domain"
)

type TaskModel struct {
	ID int
	Version int
	Title string
	Description *string
	Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time
	AuthorUserID int
}

func taskDomainsFromModels(TaskModels []TaskModel) []domain.Task {
	domains := make([]domain.Task, len(TaskModels))
	for i, model := range TaskModels {
		domains[i] = domain.NewTask(
			model.ID,
			model.Version,
			model.Title,
			model.Description,
			model.Completed,
			model.CreatedAt,
			model.CompletedAt,
			model.AuthorUserID,
		)
	}
	return domains
}