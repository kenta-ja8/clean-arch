package util

import (
	"strings"
)

func HasAcceptXml(accept string) bool {
	return strings.Contains(accept, "application/xml") || strings.Contains(accept, "text/xml")
}
