package port

import (
	"context"
	"database/sql"

	"github.com/kenta-ja8/clean-arch/pkg/entity"
)

// -- User ------
type IUserInputPort interface {
	AddUser(ctx context.Context, user *entity.User)
	GetUsers(ctx context.Context)
}

type IUserOutputPort interface {
	OutputUsers([]*entity.User)
	OutputError(error)
}

type IUserRepository interface {
	BeginTx(ctx context.Context) (*sql.Tx, error)
	AddUser(ctx context.Context, user *entity.User) ([]*entity.User, error)
	GetUsers(ctx context.Context) ([]*entity.User, error)
}

// -- Book ------
type IBookInputPort interface {
	AddBook(ctx context.Context, user *entity.Book)
	GetBooks(ctx context.Context)
}

type IBookOutputPort interface {
	OutputBooks([]*entity.Book)
	OutputError(error)
}

type IBookRepository interface {
	BeginTx(ctx context.Context) (*sql.Tx, error)
	AddBook(ctx context.Context, user *entity.Book) ([]*entity.Book, error)
	GetBooks(ctx context.Context) ([]*entity.Book, error)
}
