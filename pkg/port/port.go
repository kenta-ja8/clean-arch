package port

import (
	"context"
	"database/sql"

	"github.com/kenta-ja8/clean-arch/pkg/entity"
)

type UserInputPort interface {
	AddUser(ctx context.Context, user *entity.User) error
	GetUsers(ctx context.Context) error
}

type UserOutputPort interface {
	OutputUsers([]*entity.User) error
	OutputError(error) error
}

type UserRepository interface {
	BeginTx(ctx context.Context) (*sql.Tx, error)
	AddUser(ctx context.Context, user *entity.User) ([]*entity.User, error)
	GetUsers(ctx context.Context) ([]*entity.User, error)
}

type Datastore interface {
  Close()
}
