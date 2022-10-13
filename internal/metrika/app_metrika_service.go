package metrika

import (
	"encoding/json"
	"github.com/peepoclown/golang-yandex-metrika/internal/client"
	"github.com/peepoclown/golang-yandex-metrika/internal/request"
	"github.com/peepoclown/golang-yandex-metrika/internal/response"
	"github.com/pkg/errors"
)

type AppMetrikaService struct {
	client *client.MetrikaClient
}

type ApiTokenGetter interface {
	GetYandexApiToken() string
}

func NewAppMetrikaService(apiToken string) *AppMetrikaService {
	return &AppMetrikaService{
		client: client.NewMetrikaClient(apiToken),
	}
}

func NewAppMetrikaServiceFromConfig(apiTokenGetter ApiTokenGetter) *AppMetrikaService {
	return &AppMetrikaService{
		client: client.NewMetrikaClient(apiTokenGetter.GetYandexApiToken()),
	}
}

func (svc *AppMetrikaService) GetApplications() (*response.ApplicationsModel, error) {
	req := &request.ApplicationsRequest{}
	rawData, err := svc.client.SendAppMetrikaRequest(req)
	if err != nil {
		return nil, errors.Wrap(err, "get applications response")
	}

	responseModel := &response.ApplicationsModel{}
	err = json.Unmarshal(rawData, responseModel)
	if err != nil {
		return nil, errors.Wrap(err, "response parsing")
	}
	return responseModel, nil
}

func (svc *AppMetrikaService) GetApplication(id int64) (*response.ApplicationModel, error) {
	req := &request.ApplicationRequest{
		ApplicationId: id,
	}
	rawData, err := svc.client.SendAppMetrikaRequest(req)
	if err != nil {
		return nil, errors.Wrap(err, "get application response")
	}

	responseModel := &response.ApplicationModel{}
	err = json.Unmarshal(rawData, responseModel)
	if err != nil {
		return nil, errors.Wrap(err, "response parsing")
	}
	return responseModel, nil
}
