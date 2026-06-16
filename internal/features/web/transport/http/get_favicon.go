package web_transport_http

import (
	"net/http"

	core_logger "github.com/casuallife1334-sketch/golang-todoapp/internal/core/logger"
	core_http_response "github.com/casuallife1334-sketch/golang-todoapp/internal/core/transport/http/response"
)

func (h *WebHTTPHandler) GetFavicon(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	favicon, err := h.webService.GetFavicon()
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get favicon.svg",
		)

		return
	}

	responseHandler.FaviconResponse(favicon)
}
