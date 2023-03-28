package gateway

import (
	"context"
	"database/sql"

	"github.com/kenta-ja8/clean-arch/pkg/entity"
	"github.com/kenta-ja8/clean-arch/pkg/usecase/port"
)

type UserGateway struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) port.IUserRepository {
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
	_, err := ug.db.Exec(
		`INSERT INTO user(id, name, birthday) VALUES (?, ?, ?)`,
    user.Id,
    user.Name,
    user.Birthday,
	)
	users := []*entity.User{
    user,
	}
	return users, err
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
