package utility

import (
	"net/http"
	"strings"
)

type MethodUtility interface {
	MethodUtility()
}
type DataValue struct {
	Token   string
	Path    string
	Data    string
	Name    string
	Method  string
	payload *strings.Reader
	req     *http.Request
	err     error
	url     string
}
