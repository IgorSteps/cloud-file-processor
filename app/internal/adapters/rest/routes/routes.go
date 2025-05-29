package routes

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HandlerFactory interface {
	UploadFile() http.Handler
}

func NewRouter(handlerFactory HandlerFactory, backupLogger *slog.Logger) chi.Router {
	router := chi.NewRouter()

	router.Method(http.MethodPost, "/file", handlerFactory.UploadFile())

	return router
}
