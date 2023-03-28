package gateway

import (
	"context"
	"database/sql"

	"github.com/kenta-ja8/clean-arch/pkg/entity"
	"github.com/kenta-ja8/clean-arch/pkg/usecase/port"
)

type BookGateway struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) port.IBookRepository {
	return &BookGateway{
		db: db,
	}
}

func (ug BookGateway) BeginTx(ctx context.Context) (*sql.Tx, error) {
	tx, err := ug.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (ug BookGateway) AddBook(ctx context.Context, book *entity.Book) ([]*entity.Book, error) {
	_, err := ug.db.Exec(
		`INSERT INTO Book(id, name, birthday) VALUES (?, ?, ?)`,
    book.Id,
    book.Title,
	)
	Books := []*entity.Book{
    book,
	}
	return Books, err
}

func (ug *BookGateway) GetBooks(ctx context.Context) ([]*entity.Book, error) {
	rows, err := ug.db.Query("SELECT id, title FROM book")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*entity.Book
	for rows.Next() {
		var Book entity.Book
		if err := rows.Scan(&Book.Id, &Book.Title); err != nil {
			return nil, err
		}
		books = append(books, &Book)
	}

	return books, nil
}
