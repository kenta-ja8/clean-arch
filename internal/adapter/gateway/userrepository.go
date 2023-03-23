package gateway

import (
	"context"

	entitie "github.com/kenta-ja8/clean-arch/internal/entity"
	"github.com/kenta-ja8/clean-arch/internal/port"
)

type MySqlClientFactory interface {
	NewClient(ctx context.Context) (any, error)
}

type UserGateway struct {
	clientFactory MySqlClientFactory
}


func NewUserRepository(clientFactory MySqlClientFactory) port.UserRepository {
	return &UserGateway{
		clientFactory: clientFactory,
	}
}

func (ug UserGateway) AddUser(ctx context.Context, user *entitie.User) ([]*entitie.User, error) {
	return nil, nil
}

func (*UserGateway) GetUsers(ctx context.Context) ([]*entitie.User, error) {
	return nil, nil
}

