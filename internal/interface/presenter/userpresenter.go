package presenter

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kenta-ja8/clean-arch/internal/entity"
	"github.com/kenta-ja8/clean-arch/internal/port"
)

type UserPresenter struct {
}

func NewUserOutputPort(ctx context.Context, r http.ResponseWriter, format string) (port.UserOutputPort, error) {
	switch format {
	case "json":
		return NewJSONUserPresenter(ctx, r), nil
	case "xml":
		return NewXMLUserPresenter(ctx, r), nil
	default:
		return nil, errors.New("unsupported format")
	}
}

// func (u *UserPresenter) OutputUsers(users []*entity.User) error {
// 	return u.ctx.JSON(http.StatusOK, users)
// }
//
// func (u *UserPresenter) OutputError(err error) error {
// 	log.Fatal(err)
// 	return u.ctx.JSON(http.StatusInternalServerError, err)
// }

type UserPresenterFactory struct{}

func NewUserPresenterFactory() *UserPresenterFactory {
	return &UserPresenterFactory{}
}

type JSONUserPresenter struct {
	w http.ResponseWriter
}

func NewJSONUserPresenter(ctx context.Context, w http.ResponseWriter) *JSONUserPresenter {
	return &JSONUserPresenter{
		w: w,
	}
}

func (p *JSONUserPresenter) OutputUsers(user []*entity.User) error {
	p.w.Header().Set("Content-Type", "application/json")
	p.w.WriteHeader(http.StatusOK)
	return json.NewEncoder(p.w).Encode(user)
}
func (p *JSONUserPresenter) OutputError(err error) error {
	p.w.Header().Set("Content-Type", "application/json")
	p.w.WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(p.w).Encode(struct{}{})
}

type XMLUserPresenter struct {
	w http.ResponseWriter
}

func NewXMLUserPresenter(ctx context.Context, w http.ResponseWriter) *XMLUserPresenter {
	return &XMLUserPresenter{
		w: w,
	}
}

func (p *XMLUserPresenter) OutputUsers(user []*entity.User) error {
	p.w.Header().Set("Content-Type", "application/xml")
	p.w.WriteHeader(http.StatusOK)
	// TODO: Implement XML encoding
	return nil
}
func (p *XMLUserPresenter) OutputError(err error) error {
	p.w.Header().Set("Content-Type", "application/xml")
	p.w.WriteHeader(http.StatusBadRequest)
	// TODO: Implement XML encoding
	return nil
}
