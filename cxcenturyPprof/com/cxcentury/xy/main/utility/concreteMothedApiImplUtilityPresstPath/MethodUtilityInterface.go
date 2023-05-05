package concreteMothedApiImplUtilityPresstPath

import (
	"net/http"
	"strings"
)

type MethodUtility interface {
	MethodUtility()
}
type DataValue struct {
	Token      *string
	Data       *string
	Name       *string
	Method     *string
	payload    *strings.Reader
	req        *http.Request
	err        error
	Url        *string
	bodyString string
}

func Test() {

}
