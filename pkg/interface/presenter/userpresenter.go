package presenter

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/kenta-ja8/clean-arch/pkg/entity"
	"github.com/kenta-ja8/clean-arch/pkg/port"
	"github.com/kenta-ja8/clean-arch/pkg/util"
)

type UserPresenter struct{}

func NewUserOutputPort(ctx context.Context, r http.ResponseWriter, accept string) (port.UserOutputPort, error) {
	if util.HasAcceptXml(accept) {
		return NewXMLUserPresenter(ctx, r), nil
	}
	return NewJSONUserPresenter(ctx, r), nil
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

type Users struct {
	User []*entity.User
}
func (p *XMLUserPresenter) OutputUsers(users []*entity.User) error {
	p.w.Header().Set("Content-Type", "application/xml")
	p.w.WriteHeader(http.StatusOK)
	return xml.NewEncoder(p.w).Encode(Users{User: users})
}
func (p *XMLUserPresenter) OutputError(err error) error {
	p.w.Header().Set("Content-Type", "application/xml")
	p.w.WriteHeader(http.StatusBadRequest)
	// TODO: Implement XML encoding
	return nil
}
