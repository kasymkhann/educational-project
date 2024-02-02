package book

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	errapp "REST/internal/errApp"
	"REST/internal/handlers"
	"REST/pkg/logging"
)

var _ handlers.Handler = &handler{}

const (
	users    = "/users"
	usersURL = "/users/:id"
)

type handler struct {
	logger  *logging.Logger
	service *Service
}

func NewHandler(logger *logging.Logger, service *Service) handlers.Handler {
	return &handler{logger: logger, service: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, users, errapp.Middleware(h.CreateUserGet))
	router.HandlerFunc(http.MethodPost, users, errapp.Middleware(h.CreateUsersGet))
	router.HandlerFunc(http.MethodPut, users, errapp.Middleware(h.UpdateUsers))
	router.HandlerFunc(http.MethodPatch, users, errapp.Middleware(h.PartiallyUpdateUsers))
	router.HandlerFunc(http.MethodDelete, users, errapp.Middleware(h.UsersDelete))
}
func (h *handler) CreateUserGet(w http.ResponseWriter, r *http.Request) error {
	//TODO: реализовать логику
	return nil
}

func (h *handler) CreateUsersGet(w http.ResponseWriter, r *http.Request) error {
	//TODO: реализовать логику
	return nil
}
func (h *handler) UpdateUsers(w http.ResponseWriter, r *http.Request) error {
	//TODO: реализовать логику
	return nil
}
func (h *handler) PartiallyUpdateUsers(w http.ResponseWriter, r *http.Request) error {
	//TODO: реализовать логику
	return nil
}
func (h *handler) UsersDelete(w http.ResponseWriter, r *http.Request) error {
	//TODO: реализовать логику
	return nil
}
