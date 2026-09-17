package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mishalyamets1/golang_todoapp/internal/core/domain"
	core_logger "github.com/mishalyamets1/golang_todoapp/internal/core/logger"
	core_http_request "github.com/mishalyamets1/golang_todoapp/internal/core/transport/http/request"
	core_http_response "github.com/mishalyamets1/golang_todoapp/internal/core/transport/http/response"
)

type GetStatisticsResponse struct {
	TasksCreated int `json:"tasks_created"`
	TasksCompleted int `json:"tasks_completed"`
	TasksComletedRate *float64 `json:"tasks_completed_rate"`
	TaskAverageCompletionTime *string `json:"tasks_average_competion_time"`
}

func (h *StatisticsHTTPHandler) GetStatistics(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userId, from, to, err := getUserIdFromToQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userId/from/to query params")
		return
	}
	statisticsDomain, err := h.statisticsService.GetStatistics(ctx, userId, from, to)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get statistics")
		return
	}
	response := toDTOfromDomain(statisticsDomain)
	responseHandler.JSONResponse(response, http.StatusOK)



}

func toDTOfromDomain(statistics domain.Statistics) GetStatisticsResponse {
	var avgTime *string
	if statistics.TaskAverageCompletionTime != nil {
		duration := statistics.TaskAverageCompletionTime.String()
		avgTime = &duration
	}
	return GetStatisticsResponse{
		TasksCreated: statistics.TasksCreated,
		TasksCompleted: statistics.TasksCompleted,
		TasksComletedRate: statistics.TasksComletedRate,
		TaskAverageCompletionTime: avgTime,
	}
}

func getUserIdFromToQueryParams (r *http.Request) (*int, *time.Time, *time.Time, error) {
	const (
		fromQueryParamKey = "from"
		toQueryParamKey = "to"
		userIdQueryParamKey = "user_id"
	)

	userId, err := core_http_request.GetIntQueryParam(r, userIdQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get userId query param: %w", err)
	}
	from, err := core_http_request.GetDateQueryParam(r, fromQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get `from` query param: %w", err)
	}
	to, err := core_http_request.GetDateQueryParam(r, toQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get `to` query param: %w", err)
	}
	return userId, from, to, nil
}