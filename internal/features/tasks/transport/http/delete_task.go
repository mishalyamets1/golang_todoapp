package tasks_transport

import (
	"net/http"

	core_logger "github.com/mishalyamets1/golang_todoapp/internal/core/logger"
	core_http_request "github.com/mishalyamets1/golang_todoapp/internal/core/transport/http/request"
	core_http_response "github.com/mishalyamets1/golang_todoapp/internal/core/transport/http/response"
)

// DeleteTask godoc
// @Summary Удалить задачу
// @Description Удаление задачи по ID
// @Tags tasks
// @Param id path int true "ID удаляемой задачи"
// @Success 204 "Успешное удаление задачи"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad Request"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /tasks/{id} [delete]
func (h *TasksHTTPHandler) DeleteTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskId, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get taskId path value")
		return
	}
	if err := h.tasksService.DeleteTask(ctx, taskId); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete task")
	}
	responseHandler.NoContentResponse()
}