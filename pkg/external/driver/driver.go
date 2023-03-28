package driver

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kenta-ja8/clean-arch/pkg/adapter/controller"
	"github.com/kenta-ja8/clean-arch/pkg/adapter/gateway"
	"github.com/kenta-ja8/clean-arch/pkg/adapter/presenter"
	"github.com/kenta-ja8/clean-arch/pkg/config"
	"github.com/kenta-ja8/clean-arch/pkg/external/datastore"
	"github.com/kenta-ja8/clean-arch/pkg/usecase/interactor"
)

type IDriver interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}
type Driver struct {
	db     *sql.DB
	server *http.Server
}

func (s Driver) Start(ctx context.Context) error {
	log.Printf("Server start at %s", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s Driver) Stop(ctx context.Context) error {
	log.Println("Shutting down server...")
	err := s.server.Shutdown(ctx)
	if err != nil {
		return err
	}
	err = s.db.Close()
	return err
}

func InitializeDriver(ctx context.Context) (IDriver, error) {
	config := config.NewConfig()
	db, err := datastore.NewDataStore(ctx, config)
	if err != nil {
		return nil, err
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	userController := controller.NewUserController(
		db.DB,
		interactor.NewUserInputPort,
		presenter.NewUserOutputPort,
		gateway.NewUserRepository,
	)
	r.Get("/api/user", func(w http.ResponseWriter, r *http.Request) {
		userController.GetUsers(w, r)
	})
	r.Post("/api/user", func(w http.ResponseWriter, r *http.Request) {
		userController.AddUser(w, r)
	})

	bookController := controller.NewBookController(
		db.DB,
		interactor.NewBookInputPort,
		presenter.NewBookOutputPort,
		gateway.NewBookRepository,
	)
	r.Get("/api/book", func(w http.ResponseWriter, r *http.Request) {
		bookController.GetBooks(w, r)
	})
	r.Post("/api/book", func(w http.ResponseWriter, r *http.Request) {
		bookController.AddBook(w, r)
	})

	return Driver{
		db:     db.DB,
		server: &http.Server{Addr: config.Address, Handler: r},
	}, nil

}
