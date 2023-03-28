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

type IBookController interface {
	AddBook(w http.ResponseWriter, r *http.Request)
	GetBooks(w http.ResponseWriter, r *http.Request)
}

type BookOutputFactory func(context.Context, http.ResponseWriter, string) port.IBookOutputPort
type BookInputFactory func(port.IBookOutputPort, port.IBookRepository) port.IBookInputPort
type BookRepositoryFactory func(*sql.DB) port.IBookRepository

type BookController struct {
	db                *sql.DB
	inputFactory      BookInputFactory
	outputFactory     BookOutputFactory
	repositoryFactory BookRepositoryFactory
}

func NewBookController(
	db *sql.DB,
	inputFactory BookInputFactory,
	outputFactory BookOutputFactory,
	repositoryFactory BookRepositoryFactory,
) IBookController {
	return BookController{
		db:                db,
		inputFactory:      inputFactory,
		outputFactory:     outputFactory,
		repositoryFactory: repositoryFactory,
	}

}

func (uc BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	ip, _ := uc.newInputPort(r.Context(), w, r.Header.Get("Accept"))
	ip.GetBooks(r.Context())
}

func (uc BookController) AddBook(w http.ResponseWriter, r *http.Request) {
	ip, op := uc.newInputPort(r.Context(), w, r.Header.Get("Accept"))

	var Book *entity.Book
	err := json.NewDecoder(r.Body).Decode(&Book)
	if err != nil {
		op.OutputError(err)
		return
	}
	defer r.Body.Close()
	Book.Id = util.NewUUIDv4()

	log.Println(Book)
	validate := validator.New()
	err = validate.Struct(Book)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Println(err)
		}
		op.OutputError(err)
		return
	}

	ip.AddBook(r.Context(), Book)
}

func (u *BookController) newInputPort(c context.Context, w http.ResponseWriter, accept string) (port.IBookInputPort, port.IBookOutputPort) {
	outputPort := u.outputFactory(c, w, accept)
	repository := u.repositoryFactory(u.db)
	return u.inputFactory(outputPort, repository), outputPort
}
