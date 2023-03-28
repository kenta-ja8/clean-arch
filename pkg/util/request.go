package util

import (
	"io"
	"strings"
)

func HasAcceptXml(accept string) bool {
	return strings.Contains(accept, "application/xml") || strings.Contains(accept, "text/xml")
}

func DecodeJson(body io.ReadCloser, des interface{}) {


}
