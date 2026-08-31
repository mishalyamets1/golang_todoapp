package task_service

import (
	"context"
	"fmt"

	"github.com/mishalyamets1/golang_todoapp/internal/core/domain"
)

func (u *TasksService) GetTask(ctx context.Context, id int) (domain.Task, error) {
	task, err := u.tasksRepository.GetTask(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task from repository: %w", err)
	}
	return task, nil
}