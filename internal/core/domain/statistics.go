package domain

import "time"

type Statistics struct {
	TasksCreated               int
	TasksCompleted             int
	TaskCompletedRate          *float64
	TasksAverageCompletionTime *time.Duration
}

func NewStatistics(
	tasksCreated int,
	tasksCompleted int,
	tasksCompletedRate *float64,
	tasksAverageCompletionTime *time.Duration,
) Statistics {
	return Statistics{
		TasksCreated:               tasksCreated,
		TasksCompleted:             tasksCompleted,
		TaskCompletedRate:          tasksCompletedRate,
		TasksAverageCompletionTime: tasksAverageCompletionTime,
	}
}
