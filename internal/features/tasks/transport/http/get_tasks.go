package tasks_transport

import (
	"fmt"
	"net/http"

	core_logger "github.com/mishalyamets1/golang_todoapp/internal/core/logger"
	core_http_request "github.com/mishalyamets1/golang_todoapp/internal/core/transport/http/request"
	core_http_response "github.com/mishalyamets1/golang_todoapp/internal/core/transport/http/response"
)

type GetTasksResponse []TaskDtoResponse


func (h *TasksHTTPHandler) GetTasks(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userId, limit, offset, err := getUserIdLimitOffsetQueryParams(r)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userId/limit/offset params")
		return
	}
	taskDomains, err := h.tasksService.GetTasks(ctx, userId, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get tasks")
		return
	}
	response := GetTasksResponse(taskDTOFromDomains(taskDomains))
	responseHandler.JSONResponse(response, http.StatusOK)

}

func getUserIdLimitOffsetQueryParams(r *http.Request) (*int, *int, *int, error) {
	const (
		limitQueryParamKey = "limit"
		offsetQueryParamKey = "offset"
		userIdQueryParamKey = "user_id"
	)

	userId, err := core_http_request.GetIntQueryParam(r, "user_id")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'user_id' query param: %w", err)
	}

	limit, err := core_http_request.GetIntQueryParam(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}

	offset, err := core_http_request.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}
	return userId, limit, offset, nil
}