package consulta_cep

import (
	"encoding/json"
	"fmt"
	"github.com/giovaneboeing/desafio-cloud-run/internal/infra/service/external-http-request"
)

type RequestFactory func(method, url string) external_http_request.ExternalHttpRequestInterface

type ViaCep struct {
	HttpMethod     string
	BaseURL        string
	OutputFormat   string
	RequestFactory RequestFactory
}

func NewViaCep(requestFactory RequestFactory) *ViaCep {
	return &ViaCep{
		HttpMethod:     "GET",
		BaseURL:        "https://viacep.com.br/ws",
		OutputFormat:   "json",
		RequestFactory: requestFactory,
	}
}

func (v *ViaCep) ConsultarCep(cep string) (ConsultaCepResponse, error) {
	var viaCepResponse ConsultaCepResponse

	url := fmt.Sprintf("%s/%s/%s/", v.BaseURL, cep, v.OutputFormat)
	request := v.RequestFactory(v.HttpMethod, url)
	if err := request.IsValid(); err != nil {
		return viaCepResponse, err
	}

	response, err := request.Execute()
	if err != nil {
		return viaCepResponse, err
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return viaCepResponse, err
	}

	err = json.Unmarshal(jsonBytes, &viaCepResponse)
	if err != nil {
		return viaCepResponse, err
	}

	return viaCepResponse, nil
}
