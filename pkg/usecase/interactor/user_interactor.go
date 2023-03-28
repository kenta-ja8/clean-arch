package interactor

import (
	"context"

	"github.com/kenta-ja8/clean-arch/pkg/entity"
	"github.com/kenta-ja8/clean-arch/pkg/usecase/port"
)

type UserInteractor struct {
	OutputPort port.IUserOutputPort
	Repository port.IUserRepository
}

func NewUserInputPort(outputPort port.IUserOutputPort, repository port.IUserRepository) port.IUserInputPort {
	return &UserInteractor{
		OutputPort: outputPort,
		Repository: repository,
	}
}

func (u *UserInteractor) AddUser(ctx context.Context, user *entity.User) {
	tx, err := u.Repository.BeginTx(ctx)
	if err != nil {
		u.OutputPort.OutputError(err)
		return
	}
	defer func() {
		_ = tx.Rollback()
	}()

	users, err := u.Repository.AddUser(ctx, user)
	if err != nil {
		u.OutputPort.OutputError(err)
		return
	}
	_ = tx.Commit()

	u.OutputPort.OutputUsers(users)
}

func (u *UserInteractor) GetUsers(ctx context.Context) {
	users, err := u.Repository.GetUsers(ctx)
	if err != nil {
		u.OutputPort.OutputError(err)
		return
	}

	u.OutputPort.OutputUsers(users)
}
