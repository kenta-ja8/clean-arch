package driver

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kenta-ja8/clean-arch/pkg/adapter/controller"
	"github.com/kenta-ja8/clean-arch/pkg/adapter/gateway"
	"github.com/kenta-ja8/clean-arch/pkg/config"
	"github.com/kenta-ja8/clean-arch/pkg/datastore"
	"github.com/kenta-ja8/clean-arch/pkg/interface/presenter"
	"github.com/kenta-ja8/clean-arch/pkg/usecase/interactor"
)

type driver interface {
	Serve(ctx context.Context, address string) error
	Close()
}
type server struct {
	mux *chi.Mux
	db  *sql.DB
}

func (s server) Serve(ctx context.Context, address string) error {
	return http.ListenAndServe(address, s.mux)
}

func (s server) Close() {
	s.db.Close()
}

func InitializeDriver(ctx context.Context) (driver, error) {
	config := config.NewConfig()
	db, err := datastore.NewDataStore(ctx, config)
	if err != nil {
		return nil, err
	}

	repositoryFactory := gateway.NewUserRepository
	inputFactory := interactor.NewUserInputPort
	outputFactory := presenter.NewUserOutputPort
	controller := controller.NewUserController(db.DB, inputFactory, outputFactory, repositoryFactory)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/api/user", func(w http.ResponseWriter, r *http.Request) {
		controller.GetUsers(w, r)
	})
	r.Post("/api/user", func(w http.ResponseWriter, r *http.Request) {
		controller.AddUser(w, r)
	})

	return server{
		mux: r,
		db:  db.DB,
	}, nil

}
