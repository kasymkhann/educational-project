package users

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	errapp "REST/internal/errApp"
	"REST/internal/handlers"
	"REST/pkg/logging"
)

var _ handlers.Handler = &Handler{}

const (
	userURL  = "/users"
	usersURL = "/users/:id"
)

type Handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, userURL, errapp.Middleware(h.GetLists))
	router.HandlerFunc(http.MethodPost, userURL, errapp.Middleware(h.CreateUserPost))
	router.HandlerFunc(http.MethodGet, usersURL, errapp.Middleware(h.GetUserById))
	router.HandlerFunc(http.MethodPut, usersURL, errapp.Middleware(h.UpdateUsers))
	router.HandlerFunc(http.MethodPatch, usersURL, errapp.Middleware(h.PartiallyUpdateUsers))
	router.HandlerFunc(http.MethodDelete, usersURL, errapp.Middleware(h.UsersDelete))
}

func (h *Handler) GetLists(write http.ResponseWriter, request *http.Request) error {
	return errapp.ErrNotFound
}

func (h *Handler) CreateUserPost(write http.ResponseWriter, request *http.Request) error {
	// TODO: implement logic

	return fmt.Errorf("this is API error")
}

func (h *Handler) GetUserById(write http.ResponseWriter, request *http.Request) error {
	// TODO: implement logic
	return errapp.NewErrApp(nil, "test", "test-dev", "S-00000")
}

func (h *Handler) UpdateUsers(write http.ResponseWriter, request *http.Request) error {
	// TODO: implement logic
	write.Write([]byte("this is update users"))
	return nil
}

func (h *Handler) PartiallyUpdateUsers(writer http.ResponseWriter, request *http.Request) error {
	// TODO: implement logic
	writer.Write([]byte("this is partially users"))
	return nil

}
func (h *Handler) UsersDelete(write http.ResponseWriter, request *http.Request) error {
	// TODO: implement logic
	write.Write([]byte("this is delete users"))
	return nil
}
