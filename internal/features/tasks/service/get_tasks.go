package task_service

import (
	"context"
	"fmt"

	"github.com/mishalyamets1/golang_todoapp/internal/core/domain"
	core_errors "github.com/mishalyamets1/golang_todoapp/internal/core/errors"
)

func (s *TasksService) GetTasks(ctx context.Context, userId *int ,limit *int, offset *int) ([]domain.Task, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("limit must be non-negative: %w",core_errors.ErrInvalidArgument)
	}
	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("limit must be non-negative: %w",core_errors.ErrInvalidArgument)
	}

	tasks, err := s.tasksRepository.GetTasks(ctx, userId, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get tasks repository: %w", err)
	}
	return tasks, nil
}