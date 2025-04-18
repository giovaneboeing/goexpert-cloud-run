package external_http_request

type ExternalHttpRequestInterface interface {
	IsValid() error
	Execute() (interface{}, error)
}
