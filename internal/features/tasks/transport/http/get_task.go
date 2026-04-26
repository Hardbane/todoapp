package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/Hardbane/todoapp/internal/core/logger"
	core_http_request "github.com/Hardbane/todoapp/internal/core/transport/http/request"
	core_http_response "github.com/Hardbane/todoapp/internal/core/transport/http/response"
)

type GetTaskResponse TaskDTOResponse

// GetTask godoc
// @Summary Получить конкретную задачу
// @Tags tasks
// @Produce json
// @Param id path int true "ID получаемой задачи"
// @Success 200 {object} GetTaskResponse "Задача успешно найдена"
// @Failure 400 {object} core_http_response.ErrorResponse "bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "task not found"
// @Failure 500 {object} core_http_response.ErrorResponse "server error"
// @Router /tasks/{id} [get]
func (h *TasksHTTPHandler) GetTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get taskID path value",
		)
		return
	}

	taskDomain, err := h.tasksService.GetTask(ctx, taskID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get task",
		)
		return
	}
	response := GetTaskResponse(taskDTOFromDomain(taskDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}
