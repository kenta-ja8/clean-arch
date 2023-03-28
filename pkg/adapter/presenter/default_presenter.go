package presenter

import (
	"encoding/json"
	"net/http"
)

// -- JSON ------
type DefaultJSONPresenter struct {
	w http.ResponseWriter
}
type ErrorOutput struct {
	ErrorMsg string `json:"errorMsg"`
}

func (p DefaultJSONPresenter) OutputError(err error) {
	p.w.Header().Set("Content-Type", "application/json")
	p.w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(p.w).Encode(ErrorOutput{
		ErrorMsg: err.Error(),
	})
}

// -- XML ------
type DefaultXmlPresenter struct {
	w http.ResponseWriter
}

func (p *DefaultXmlPresenter) OutputError(err error) {
	p.w.Header().Set("Content-Type", "application/xml")
	p.w.WriteHeader(http.StatusBadRequest)
	// TODO: Implement XML encoding
}
