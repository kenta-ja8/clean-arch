package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kenta-ja8/clean-arch/pkg/entity"
	"github.com/kenta-ja8/clean-arch/pkg/usecase/port"
	"github.com/kenta-ja8/clean-arch/pkg/util"
)

type IUserController interface {
	AddUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
}

type UserOutputFactory func(context.Context, http.ResponseWriter, string) port.IUserOutputPort
type UserInputFactory func(port.IUserOutputPort, port.IUserRepository) port.IUserInputPort
type UserRepositoryFactory func(*sql.DB) port.IUserRepository

type UserController struct {
	db                *sql.DB
	inputFactory      UserInputFactory
	outputFactory     UserOutputFactory
	repositoryFactory UserRepositoryFactory
}

func NewUserController(
	db *sql.DB,
	inputFactory UserInputFactory,
	outputFactory UserOutputFactory,
	repositoryFactory UserRepositoryFactory,
) IUserController {
	return UserController{
		db:                db,
		inputFactory:      inputFactory,
		outputFactory:     outputFactory,
		repositoryFactory: repositoryFactory,
	}

}

func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	ip, _ := uc.newInputPort(r.Context(), w, r.Header.Get("Accept"))
	ip.GetUsers(r.Context())
}

func (uc UserController) AddUser(w http.ResponseWriter, r *http.Request) {
	ip, op := uc.newInputPort(r.Context(), w, r.Header.Get("Accept"))

	var user *entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		op.OutputError(err)
		return
	}
	defer r.Body.Close()
	user.Id = util.NewUUIDv4()

	log.Println(user)
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Println(err)
		}
		op.OutputError(err)
		return
	}

	ip.AddUser(r.Context(), user)
}

func (u *UserController) newInputPort(c context.Context, w http.ResponseWriter, accept string) (port.IUserInputPort, port.IUserOutputPort) {
	outputPort := u.outputFactory(c, w, accept)
	repository := u.repositoryFactory(u.db)
	return u.inputFactory(outputPort, repository), outputPort
}
