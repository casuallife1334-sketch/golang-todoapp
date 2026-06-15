package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/casuallife1334-sketch/golang-todoapp/internal/core/logger"
	core_http_request "github.com/casuallife1334-sketch/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/casuallife1334-sketch/golang-todoapp/internal/core/transport/http/response"
)

// DeleteTask 	godoc
// @Summary 	Удаление задачи
// @Description Удаление существующей задачи в системе по её ID
// @Tags 		tasks
// @Param 		id path int true "ID удаляемой задачи"
// @Success 	204 "Успешное удаление задачи"
// @Failure 	400 {object} core_http_response.ErrorResponse "Bad Request"
// @Failure 	404 {object} core_http_response.ErrorResponse "task not found"
// @Failure 	500 {object} core_http_response.ErrorResponse "internal server error"
// @Router 		/tasks/{id} [delete]
func (h *TasksHTTPHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get taskID path value",
		)

		return
	}

	if err := h.taskService.DeleteTask(ctx, taskID); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete task",
		)

		return
	}

	responseHandler.NoContentResponse()
}
