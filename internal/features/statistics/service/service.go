package statistics_service

import (
	"context"
	"time"

	"github.com/casuallife1334-sketch/golang-todoapp/internal/core/domain"
)

type StatisticsRepository interface {
	GetTasks(
		ctx context.Context,
		userID *int,
		from *time.Time,
		to *time.Time,
	) ([]domain.Task, error)
}

type StatisticsService struct {
	statisticsRepository StatisticsRepository
}

func NewStatisticsService(
	statisticsRepository StatisticsRepository,
) *StatisticsService {
	return &StatisticsService{
		statisticsRepository: statisticsRepository,
	}
}
