package presenter

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"time"

	"github.com/kenta-ja8/clean-arch/pkg/entity"
	"github.com/kenta-ja8/clean-arch/pkg/usecase/port"
	"github.com/kenta-ja8/clean-arch/pkg/util"
)

func NewUserOutputPort(ctx context.Context, r http.ResponseWriter, accept string) port.IUserOutputPort {
	log.Println("Accept Header: ", accept)
	if util.HasAcceptXml(accept) {
		return NewXMLUserPresenter(ctx, r)
	}
	return NewJSONUserPresenter(ctx, r)
}

// -- JSON ------
func NewJSONUserPresenter(ctx context.Context, w http.ResponseWriter) *JsonUserPresenter {
	return &JsonUserPresenter{
		DefaultJSONPresenter{
			w: w,
		},
	}
}

type JsonUserPresenter struct {
	DefaultJSONPresenter
}

type UserOutput struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
	Age      int       `json:"age"`
}

func (p *JsonUserPresenter) OutputUsers(users []*entity.User) {
	p.w.Header().Set("Content-Type", "application/json")
	p.w.WriteHeader(http.StatusOK)
	output := make([]*UserOutput, 0, len(users))
	for _, user := range users {
		output = append(output, &UserOutput{
			Id:       user.Id,
			Name:     user.Name,
			Birthday: user.Birthday,
			Age:      user.Age(),
		})
	}
	_ = json.NewEncoder(p.w).Encode(output)
}

// -- XML ------
type XmlUserPresenter struct {
	DefaultXmlPresenter
}

func NewXMLUserPresenter(ctx context.Context, w http.ResponseWriter) *XmlUserPresenter {
	return &XmlUserPresenter{
		DefaultXmlPresenter{
			w: w,
		},
	}
}

type Users struct {
	User []*entity.User
}

func (p *XmlUserPresenter) OutputUsers(users []*entity.User) {
	p.w.Header().Set("Content-Type", "application/xml")
	p.w.WriteHeader(http.StatusOK)
	_ = xml.NewEncoder(p.w).Encode(Users{User: users})
}
