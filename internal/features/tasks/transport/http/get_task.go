package tasks_transport

import (
	"net/http"

	core_logger "github.com/mishalyamets1/golang_todoapp/internal/core/logger"
	core_http_request "github.com/mishalyamets1/golang_todoapp/internal/core/transport/http/request"
	core_http_response "github.com/mishalyamets1/golang_todoapp/internal/core/transport/http/response"
)
type GetTaskResponse TaskDtoResponse

// GetTask godoc
// @Summary Получить задачу
// @Description Получить задачу по ID
// @Tags tasks
// @Produce json
// @Param id path int true "ID задачи"
// @Success 200 {object} GetTaskResponse "Задача"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad Request"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /tasks/{id} [get]
func (h *TasksHTTPHandler) GetTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskId, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get path value")
		return
	}
	taskDomain, err := h.tasksService.GetTask(ctx, taskId)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task")
		return
	}
	response := GetTaskResponse(taskDTOFromDomain(taskDomain))
	responseHandler.JSONResponse(response, http.StatusOK)
}