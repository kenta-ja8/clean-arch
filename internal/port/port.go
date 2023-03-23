package port

import (
	"context"

	"github.com/kenta-ja8/clean-arch/internal/entity"
)

// 1. ユーザーデータを受け取り
type UserInputPort interface {
	AddUser(ctx context.Context, user *entity.User) error
	GetUsers(ctx context.Context) error
}

// 3. 全てのユーザーを返却 or エラーを返す
type UserOutputPort interface {
	OutputUsers([]*entity.User) error
	OutputError(error) error
}

// 2. 何かしらにユーザーデータを追加、全ユーザーを返す
type UserRepository interface {
	AddUser(ctx context.Context, user *entity.User) ([]*entity.User, error)
	GetUsers(ctx context.Context) ([]*entity.User, error)
}
