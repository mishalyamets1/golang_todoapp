package task_service

import (
	"context"
	"fmt"

	"github.com/mishalyamets1/golang_todoapp/internal/core/domain"
)

func (s *TasksService) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	if err := task.Validate(); err != nil {
		return domain.Task{}, fmt.Errorf("validate domain task: %w", err)
	}
	task, err := s.tasksRepository.CreateTask(ctx, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("create task: %w", err)
	}
	return task, nil
}