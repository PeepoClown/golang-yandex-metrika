package request

type MetrikaRequest interface {
	GetPath() string
	GetRequestParams() map[string]string
}
