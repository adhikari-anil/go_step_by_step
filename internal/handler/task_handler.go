package handler

import (
	"fmt"
	"net/http"
)

type TaskHandler struct{}

func NewTaskHandlers() *TaskHandler {
	return &TaskHandler{}
}

func (h *TaskHandler) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is actual use of handlers!")
}

func (h *TaskHandler) Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK!")
}
