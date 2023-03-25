package gateway

import (
	"context"
	"database/sql"

	"github.com/kenta-ja8/clean-arch/pkg/datastore"
	"github.com/kenta-ja8/clean-arch/pkg/entity"
	"github.com/kenta-ja8/clean-arch/pkg/port"
)

type DatastoreFactory interface {
	NewDatastore(ctx context.Context) (datastore.Datastore, error)
}

type UserGateway struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) port.UserRepository {
	return &UserGateway{
		db: db,
	}
}

func (ug UserGateway) BeginTx(ctx context.Context) (*sql.Tx, error) {
	tx, err := ug.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (ug UserGateway) AddUser(ctx context.Context, user *entity.User) ([]*entity.User, error) {

	return nil, nil
}

func (ug *UserGateway) GetUsers(ctx context.Context) ([]*entity.User, error) {
	rows, err := ug.db.Query("SELECT id, name, birthday FROM user")
	if err != nil {
    return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Birthday); err != nil {
      return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}
