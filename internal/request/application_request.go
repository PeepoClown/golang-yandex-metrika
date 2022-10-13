package request

import (
	"fmt"
	"github.com/peepoclown/golang-yandex-metrika/internal/constants"
)

type ApplicationRequest struct {
	ApplicationId int64
	path          string
}

func (req *ApplicationRequest) GetPath() string {
	return fmt.Sprintf(constants.AppMetrikaApiApplicationPath, req.ApplicationId)
}

func (req *ApplicationRequest) GetRequestParams() map[string]string {
	return make(map[string]string)
}
