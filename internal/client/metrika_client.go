package client

import (
	"fmt"
	"github.com/peepoclown/golang-yandex-metrika/internal/constants"
	"github.com/peepoclown/golang-yandex-metrika/internal/request"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"time"
)

type MetrikaClient struct {
	httpClient *http.Client
	apiToken   string
}

const (
	DefaultTimeout = time.Duration(60) * time.Second
)

func NewMetrikaClient(apiToken string) *MetrikaClient {
	return &MetrikaClient{
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		apiToken: apiToken,
	}
}

func (metrikaClient *MetrikaClient) SendAppMetrikaRequest(requestModel request.MetrikaRequest) ([]byte, error) {
	return metrikaClient.SendRequest(
		fmt.Sprintf("%s/%s", constants.AppMetrikaApiBaseUrl, requestModel.GetPath()),
		requestModel.GetRequestParams(),
	)
}

func (metrikaClient *MetrikaClient) SendRequest(url string, params map[string]string) ([]byte, error) {
	oAuthToken := "OAuth " + metrikaClient.apiToken

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", oAuthToken)
	req.Header.Add("Content-Type", "application/json")

	queryParams := req.URL.Query()
	for k, v := range params {
		queryParams.Add(k, v)
	}
	req.URL.RawQuery = queryParams.Encode()

	if err != nil {
		return nil, errors.Wrap(err, "http request creation")
	}

	resp, err := metrikaClient.httpClient.Do(req)
	if err != nil || resp.Body == nil {
		if err == nil {
			return nil, errors.New("empty response body")
		}
		return nil, errors.Wrap(err, "http request sending")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "http response processing")
	}

	return body, nil
}
