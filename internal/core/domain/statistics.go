package domain

import "time"

type Statistics struct {
	TasksCreated int
	TasksCompleted int
	TasksComletedRate *float64
	TaskAverageCompletionTime *time.Duration
}

func NewStatistics(tasksCreated int, tasksCompleted int, tasksComletedRate *float64, taskAverageCompletionTime *time.Duration) Statistics {
	return Statistics{
		TasksCreated: tasksCreated,
		TasksCompleted: tasksCompleted,
		TasksComletedRate: tasksComletedRate,
		TaskAverageCompletionTime: taskAverageCompletionTime,
	}
}