package interactor

import (
	"context"

	"github.com/kenta-ja8/clean-arch/internal/entity"
	"github.com/kenta-ja8/clean-arch/internal/port"
)

type UserInteractor struct {
	OutputPort port.UserOutputPort
	Repository port.UserRepository
}

func NewUserInputPort(outputPort port.UserOutputPort, repository port.UserRepository) port.UserInputPort {
	return &UserInteractor{
		OutputPort: outputPort,
		Repository: repository,
	}
}

func (u *UserInteractor) AddUser(ctx context.Context, user *entity.User) error {
	users, err := u.Repository.AddUser(ctx, user)
	if err != nil {
		return u.OutputPort.OutputError(err)
	}

	return u.OutputPort.OutputUsers(users)
}

func (u *UserInteractor) GetUsers(ctx context.Context) error {
	users, err := u.Repository.GetUsers(ctx)
	if err != nil {
		return u.OutputPort.OutputError(err)
	}

	return u.OutputPort.OutputUsers(users)
}
