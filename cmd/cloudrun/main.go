package main

import (
	"github.com/giovaneboeing/desafio-cloud-run/configs"
	consulta_cep "github.com/giovaneboeing/desafio-cloud-run/internal/infra/service/consulta-cep"
	consulta_temperatura "github.com/giovaneboeing/desafio-cloud-run/internal/infra/service/consulta-temperatura"
	external_http_request "github.com/giovaneboeing/desafio-cloud-run/internal/infra/service/external-http-request"
	"github.com/giovaneboeing/desafio-cloud-run/internal/infra/usecase"
	"github.com/giovaneboeing/desafio-cloud-run/internal/infra/web/webserver"
	"github.com/giovaneboeing/desafio-cloud-run/internal/infra/web/webserver/handler"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	consultaCepService := consulta_cep.NewViaCep(externalHttpRequestFactory)
	consultaTemperaturaService := consulta_temperatura.NewWeatherApi(cfg.WeatherApiKey, externalHttpRequestFactory)
	temperaturaUseCase := usecase.NewConsultaTemperaturaUseCase(consultaCepService, consultaTemperaturaService)

	temperaturaHandler := handler.NewTemperaturaHandler(temperaturaUseCase)

	webserver := webserver.NewWebServer(cfg.WebServerPort)
	webserver.AddHandler("/temperatura", temperaturaHandler.Consultar)
	webserver.Start()
}

func externalHttpRequestFactory(method, url string) external_http_request.ExternalHttpRequestInterface {
	return external_http_request.NewHttpRequest(method, url)
}
