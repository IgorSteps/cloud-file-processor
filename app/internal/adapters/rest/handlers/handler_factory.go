package handlers

import "net/http"

type HandlerFactory struct {
	uploadFile *UploadFile
}

func (s *HandlerFactory) UploadFile() http.Handler {
	return s.uploadFile
}
