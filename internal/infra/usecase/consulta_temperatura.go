package usecase

import (
	"errors"
	"github.com/giovaneboeing/desafio-cloud-run/internal/infra/service/consulta-cep"
	"github.com/giovaneboeing/desafio-cloud-run/internal/infra/service/consulta-temperatura"
)

type ConsultaTemperaturaOutputDto struct {
	Cep                   string  `json:"cep,omitempty"`
	Cidade                string  `json:"cidade,omitempty"`
	TemperaturaCelcius    float64 `json:"temp_C,omitempty"`
	TemperaturaFahrenheit float64 `json:"temp_F,omitempty"`
	TemperaturaKelvin     float64 `json:"temp_K,omitempty"`
}

func NewConsultaTemperaturaOutputDto(cep, cidade string, temperaturaCelcius float64) ConsultaTemperaturaOutputDto {
	return ConsultaTemperaturaOutputDto{
		Cep:                   cep,
		Cidade:                cidade,
		TemperaturaCelcius:    temperaturaCelcius,
		TemperaturaFahrenheit: (temperaturaCelcius * 9 / 5) + 32,
		TemperaturaKelvin:     temperaturaCelcius + 273.15,
	}
}

type ConsultaTemperaturaUseCase struct {
	ConsultaCep         consulta_cep.CepServiceInterface
	ConsultaTemperatura consulta_temperatura.TemperaturaServiceInterface
}

func NewConsultaTemperaturaUseCase(consultaCep consulta_cep.CepServiceInterface, consultaTemperatura consulta_temperatura.TemperaturaServiceInterface) *ConsultaTemperaturaUseCase {
	return &ConsultaTemperaturaUseCase{
		ConsultaCep:         consultaCep,
		ConsultaTemperatura: consultaTemperatura,
	}
}

func (c *ConsultaTemperaturaUseCase) Execute(cep string) (ConsultaTemperaturaOutputDto, error) {
	consultaCepResponse, err := c.ConsultaCep.ConsultarCep(cep)
	if err != nil {
		return ConsultaTemperaturaOutputDto{}, err
	}

	if consultaCepResponse.Erro == "true" {
		return ConsultaTemperaturaOutputDto{}, errors.New("can not find zipcode")
	}

	consultaTemperaturaResponse, err := c.ConsultaTemperatura.ConsultaTemperatura(consultaCepResponse.Cidade)
	if err != nil {
		return ConsultaTemperaturaOutputDto{}, err
	}

	return NewConsultaTemperaturaOutputDto(consultaCepResponse.Cep, consultaCepResponse.Cidade, consultaTemperaturaResponse.Current.TempC), nil
}
