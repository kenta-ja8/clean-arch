package presenter

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"

	"github.com/kenta-ja8/clean-arch/pkg/entity"
	"github.com/kenta-ja8/clean-arch/pkg/usecase/port"
	"github.com/kenta-ja8/clean-arch/pkg/util"
)

func NewBookOutputPort(ctx context.Context, r http.ResponseWriter, accept string) port.IBookOutputPort {
	log.Println("Accept Header: ", accept)
	if util.HasAcceptXml(accept) {
		return NewXMLBookPresenter(ctx, r)
	}
	return NewJSONBookPresenter(ctx, r)
}

// -- JSON ------
func NewJSONBookPresenter(ctx context.Context, w http.ResponseWriter) *JsonBookPresenter {
	return &JsonBookPresenter{
		DefaultJSONPresenter{
			w: w,
		},
	}
}

type JsonBookPresenter struct {
	DefaultJSONPresenter
}

type BookOutput struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

func (p *JsonBookPresenter) OutputBooks(Books []*entity.Book) {
	p.w.Header().Set("Content-Type", "application/json")
	p.w.WriteHeader(http.StatusOK)
	output := make([]*BookOutput, 0, len(Books))
	for _, Book := range Books {
		output = append(output, &BookOutput{
			Id:    Book.Id,
			Title: Book.Title,
		})
	}
	_ = json.NewEncoder(p.w).Encode(output)
}

// -- XML ------
type XmlBookPresenter struct {
	DefaultXmlPresenter
}

func NewXMLBookPresenter(ctx context.Context, w http.ResponseWriter) *XmlBookPresenter {
	return &XmlBookPresenter{
		DefaultXmlPresenter{
			w: w,
		},
	}
}

type Books struct {
	Book []*entity.Book
}

func (p *XmlBookPresenter) OutputBooks(books []*entity.Book) {
	p.w.Header().Set("Content-Type", "application/xml")
	p.w.WriteHeader(http.StatusOK)
	_ = xml.NewEncoder(p.w).Encode(Books{Book: books})
}
