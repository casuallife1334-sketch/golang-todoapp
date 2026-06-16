package web_transport_http

import (
	core_http_server "github.com/casuallife1334-sketch/golang-todoapp/internal/core/transport/http/server"
)

type WebService interface {
	GetMainPage() ([]byte, error)
	GetFavicon() ([]byte, error)
}

type WebHTTPHandler struct {
	webService WebService
}

func NewWebHTTPHandler(webService WebService) *WebHTTPHandler {
	return &WebHTTPHandler{
		webService: webService,
	}
}

func (h *WebHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Path:    "/",
			Handler: h.GetMainPage,
		},
		{
			Path:    "/favicon.svg",
			Handler: h.GetFavicon,
		},
	}
}
