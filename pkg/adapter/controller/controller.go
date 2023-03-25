package controller

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/kenta-ja8/clean-arch/pkg/entity"
	"github.com/kenta-ja8/clean-arch/pkg/port"
)

type User interface {
	AddUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
}

type OutputFactory func(context.Context, http.ResponseWriter, string) (port.UserOutputPort, error)
type InputFactory func(port.UserOutputPort, port.UserRepository) port.UserInputPort
type RepositoryFactory func(*sql.DB) port.UserRepository
type DatastoreFactory func() port.Datastore

type UserController struct {
	db                *sql.DB
	inputFactory      InputFactory
	outputFactory     OutputFactory
	repositoryFactory RepositoryFactory
}

func NewUserController(
	db *sql.DB,
	inputFactory InputFactory,
	outputFactory OutputFactory,
	repositoryFactory RepositoryFactory,
) User {
	return UserController{
		db:                db,
		inputFactory:      inputFactory,
		outputFactory:     outputFactory,
		repositoryFactory: repositoryFactory,
	}

}

func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	_ = uc.newInputPort(r.Context(), w, r.Header.Get("Accept")).GetUsers(r.Context())
}

func (uc UserController) AddUser(w http.ResponseWriter, r *http.Request) {
	var user *entity.User
	_ = uc.newInputPort(r.Context(), w, r.Header.Get("Accept")).AddUser(r.Context(), user)

}

func (u *UserController) newInputPort(c context.Context, w http.ResponseWriter, format string) port.UserInputPort {
	outputPort, _ := u.outputFactory(c, w, format)
	repository := u.repositoryFactory(u.db)
	return u.inputFactory(outputPort, repository)
}
