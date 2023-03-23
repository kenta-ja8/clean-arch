package controller

import (
	"context"
	"net/http"

	"github.com/kenta-ja8/clean-arch/internal/adapter/gateway"
	"github.com/kenta-ja8/clean-arch/internal/port"
)

type User interface {
	AddUser(ctx context.Context) http.HandlerFunc
	GetUsers(ctx context.Context) http.HandlerFunc
}

type OutputFactory func(context.Context, http.ResponseWriter, string) (port.UserOutputPort,error)
type InputFactory func(port.UserOutputPort, port.UserRepository) port.UserInputPort
type RepositoryFactory func(gateway.MySqlClientFactory) port.UserRepository

type UserController struct {
	inputFactory      InputFactory
	outputFactory     OutputFactory
	repositoryFactory RepositoryFactory
  clientFactory     gateway.MySqlClientFactory
}

func NewUserController(
	inputFactory InputFactory,
	outputFactory OutputFactory,
	repositoryFactory RepositoryFactory,
	clientFactory gateway.MySqlClientFactory,
) User {
	return UserController{
		inputFactory:      inputFactory,
		outputFactory:     outputFactory,
		repositoryFactory: repositoryFactory,
		clientFactory:     clientFactory,
	}

}

func (uc UserController) GetUsers(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// reqres.JSON(200, struct{ id string }{id: "xxx"})
		// reqres.String(200, "pong")
		_ = uc.newInputPort(r.Context(), w, "json").GetUsers(r.Context())

		// ctx, cancel := context.WithCancel(ginCtx)
	}
}

func (uc UserController) AddUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (u *UserController) newInputPort(c context.Context, w http.ResponseWriter, format string) port.UserInputPort {
	outputPort, _ := u.outputFactory(c, w, format)
	repository := u.repositoryFactory(u.clientFactory)
	return u.inputFactory(outputPort, repository)
}
