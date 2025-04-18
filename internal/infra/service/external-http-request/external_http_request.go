package external_http_request

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ExternalHttpRequest struct {
	Method string
	Url    string
}

func NewHttpRequest(method, url string) *ExternalHttpRequest {
	return &ExternalHttpRequest{
		Method: method,
		Url:    url,
	}
}

func (h *ExternalHttpRequest) IsValid() error {
	if h.Method == "" {
		return errors.New("invalid method")
	}
	if h.Url == "" {
		return errors.New("invalid url")
	}
	return nil
}

func (h *ExternalHttpRequest) Execute() (interface{}, error) {
	if err := h.IsValid(); err != nil {
		return nil, err
	}

	httpClient := http.Client{}
	req, err := http.NewRequest(h.Method, h.Url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to execute request")
	}

	var responseRequest interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseRequest); err != nil {
		return nil, err
	}
	return responseRequest, nil
}
