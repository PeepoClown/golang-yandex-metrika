package request

import (
	"github.com/peepoclown/golang-yandex-metrika/internal/constants"
)

type ApplicationsRequest struct {
	path string
}

func (req *ApplicationsRequest) GetPath() string {
	return constants.AppMetrikaApiApplicationsPath
}

func (req *ApplicationsRequest) GetRequestParams() map[string]string {
	return make(map[string]string)
}
