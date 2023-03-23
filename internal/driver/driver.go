package driver

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kenta-ja8/clean-arch/internal/adapter/controller"
	"github.com/kenta-ja8/clean-arch/internal/adapter/gateway"
	"github.com/kenta-ja8/clean-arch/internal/database"
	"github.com/kenta-ja8/clean-arch/internal/interface/presenter"
	"github.com/kenta-ja8/clean-arch/internal/usecase/interactor"
)

type driver interface {
	Serve(ctx context.Context, address string) error
}
type server struct {
	mux *chi.Mux
}

func (s server) Serve(ctx context.Context, address string) error {
	return http.ListenAndServe(address, s.mux)
}

func InitializeDriver(ctx context.Context) driver {
	// mux := http.NewServeMux()
	// r := gin.Default()
	inputFactory := NewInputFactory()
	outputFactory := NewOutputFactory()
	repositoryFactory := NewRepositoryFactory()
	mySqlClientFactory := NewMySqlClientFactory()

	controller := controller.NewUserController(inputFactory, outputFactory, repositoryFactory, mySqlClientFactory)
	// controller.AddUser(ctx)
	//
	// r.GET("/api/user", func(ctx *gin.Context) {
	// 	controller.GetUsers(ctx)(ctx, ctx)
	// })
	// r.POST("/api/user", func(ctx *gin.Context) {
	// 	controller.AddUser(ctx)
	// })

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		controller.GetUsers(ctx)(w, r)
	})

	return server{
		mux: r,
	}

}

func NewOutputFactory() controller.OutputFactory {
	return presenter.NewUserOutputPort
}

func NewInputFactory() controller.InputFactory {
	return interactor.NewUserInputPort
}

func NewRepositoryFactory() controller.RepositoryFactory {
	return gateway.NewUserRepository
}
func NewMySqlClientFactory() gateway.MySqlClientFactory {
	return &database.MySqlClientFactory{}
}
