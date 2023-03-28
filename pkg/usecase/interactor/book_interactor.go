package interactor

import (
	"context"

	"github.com/kenta-ja8/clean-arch/pkg/entity"
	"github.com/kenta-ja8/clean-arch/pkg/usecase/port"
)

type BookInteractor struct {
	OutputPort port.IBookOutputPort
	Repository port.IBookRepository
}

func NewBookInputPort(outputPort port.IBookOutputPort, repository port.IBookRepository) port.IBookInputPort {
	return &BookInteractor{
		OutputPort: outputPort,
		Repository: repository,
	}
}

func (u *BookInteractor) AddBook(ctx context.Context, book *entity.Book) {
	tx, err := u.Repository.BeginTx(ctx)
	if err != nil {
		u.OutputPort.OutputError(err)
		return
	}
	defer func() {
		_ = tx.Rollback()
	}()

	books, err := u.Repository.AddBook(ctx, book)
	if err != nil {
		u.OutputPort.OutputError(err)
		return
	}
	_ = tx.Commit()

	u.OutputPort.OutputBooks(books)
}

func (u *BookInteractor) GetBooks(ctx context.Context) {
	books, err := u.Repository.GetBooks(ctx)
	if err != nil {
		u.OutputPort.OutputError(err)
		return
	}

	u.OutputPort.OutputBooks(books)
}
