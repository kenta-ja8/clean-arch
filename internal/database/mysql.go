package database

import (
	"context"
)

type MySqlClientFactory struct{}

func (f *MySqlClientFactory) NewClient(ctx context.Context) (any, error) {
	return nil, nil
}
