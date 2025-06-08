package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HandlerFactory interface {
	UploadFile() http.Handler
}

func NewRouter(handlerFactory HandlerFactory, middlewares []func(http.Handler) http.Handler) chi.Router {
	router := chi.NewRouter()

	for _, mw := range middlewares {
		router.Use(mw)
	}

	router.Method(http.MethodPost, "/file", handlerFactory.UploadFile())

	return router
}
