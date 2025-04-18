package consulta_temperatura

import (
	"encoding/json"
	"fmt"
	"github.com/giovaneboeing/desafio-cloud-run/internal/infra/service/external-http-request"
	"golang.org/x/text/unicode/norm"
	"net/url"
	"strings"
	"unicode"
)

type RequestFactory func(method, url string) external_http_request.ExternalHttpRequestInterface

type WeatherApi struct {
	ApiKey         string
	HttpMethod     string
	BaseURL        string
	RequestFactory RequestFactory
}

func NewWeatherApi(apiKey string, requestFactory RequestFactory) *WeatherApi {
	return &WeatherApi{
		ApiKey:         apiKey,
		BaseURL:        "http://api.weatherapi.com/v1/current.json",
		HttpMethod:     "GET",
		RequestFactory: requestFactory,
	}
}

func (w *WeatherApi) ConsultaTemperatura(cidade string) (ConsultaTemperaturaResponse, error) {
	var weatherAPIResponse ConsultaTemperaturaResponse

	encodedCidade := w.codificarNomeCidade(cidade)
	url := fmt.Sprintf("%s?key=%s&q=%s&aqi=no", w.BaseURL, w.ApiKey, encodedCidade)
	request := w.RequestFactory(w.HttpMethod, url)
	if err := request.IsValid(); err != nil {
		return weatherAPIResponse, err
	}

	response, err := request.Execute()
	if err != nil {
		return weatherAPIResponse, err
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return weatherAPIResponse, err
	}

	err = json.Unmarshal(jsonBytes, &weatherAPIResponse)
	if err != nil {
		return weatherAPIResponse, err
	}

	return weatherAPIResponse, nil
}

func (w *WeatherApi) codificarNomeCidade(cidade string) interface{} {
	t := norm.NFD.String(cidade)

	var b strings.Builder
	b.Grow(len(t))

	for _, r := range t {
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		b.WriteRune(r)
	}
	return url.PathEscape(b.String())
}
