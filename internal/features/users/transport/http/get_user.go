package user_transport_http

import (
	"net/http"

	core_logger "github.com/mishalyamets1/golang_todoapp/internal/core/logger"
	core_http_request "github.com/mishalyamets1/golang_todoapp/internal/core/transport/http/request"
	core_http_response "github.com/mishalyamets1/golang_todoapp/internal/core/transport/http/response"
)

type GetUserResponse UserDTOResponse

// GetUser godoc
// @Summary Получить пользователя
// @Description Получить пользователя по ID
// @Tags users
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} GetUserResponse "Пользователь"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad Request"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /users/{id} [get]
func (h *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request)  {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userId, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get path value")
		return
	}
	user, err := h.usersService.GetUser(ctx, userId)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
		return
	}

	response := GetUserResponse(userDTOFromDomain(user))
	responseHandler.JSONResponse(response, http.StatusOK)
}