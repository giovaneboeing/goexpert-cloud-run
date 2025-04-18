package usecase

import (
	"errors"
	consulta_cep "github.com/giovaneboeing/desafio-cloud-run/internal/infra/service/consulta-cep"
	consulta_temperatura "github.com/giovaneboeing/desafio-cloud-run/internal/infra/service/consulta-temperatura"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockCepService struct {
	mock.Mock
}

func (m *mockCepService) ConsultarCep(cep string) (consulta_cep.ConsultaCepResponse, error) {
	args := m.Called(cep)
	return args.Get(0).(consulta_cep.ConsultaCepResponse), args.Error(1)
}

type mockTemperaturaService struct {
	mock.Mock
}

func (m *mockTemperaturaService) ConsultaTemperatura(cidade string) (consulta_temperatura.ConsultaTemperaturaResponse, error) {
	args := m.Called(cidade)
	return args.Get(0).(consulta_temperatura.ConsultaTemperaturaResponse), args.Error(1)
}

func TestExecute_Success(t *testing.T) {
	mockCep := new(mockCepService)
	mockTemp := new(mockTemperaturaService)

	cep := "88000000"
	cidade := "Florianópolis"
	tempC := 25.0

	mockCepResponse := consulta_cep.ConsultaCepResponse{
		Cep:    cep,
		Cidade: cidade,
	}

	mockTempResponse := consulta_temperatura.ConsultaTemperaturaResponse{}
	mockTempResponse.Current.TempC = tempC

	mockCep.On("ConsultarCep", cep).Return(mockCepResponse, nil)
	mockTemp.On("ConsultaTemperatura", cidade).Return(mockTempResponse, nil)

	useCase := NewConsultaTemperaturaUseCase(mockCep, mockTemp)

	result, err := useCase.Execute(cep)

	assert.NoError(t, err)
	assert.Equal(t, cep, result.Cep)
	assert.Equal(t, cidade, result.Cidade)
	assert.Equal(t, 25.0, result.TemperaturaCelcius)
	assert.Equal(t, 77.0, result.TemperaturaFahrenheit)
	assert.Equal(t, 298.15, result.TemperaturaKelvin)

	mockCep.AssertExpectations(t)
	mockTemp.AssertExpectations(t)
}

func TestExecute_ErroConsultaCep(t *testing.T) {
	mockCep := new(mockCepService)
	mockTemp := new(mockTemperaturaService)

	cep := "88000000"
	mockCep.On("ConsultarCep", cep).Return(consulta_cep.ConsultaCepResponse{}, errors.New("erro no serviço de CEP"))

	useCase := NewConsultaTemperaturaUseCase(mockCep, mockTemp)

	result, err := useCase.Execute(cep)

	assert.Error(t, err)
	assert.EqualError(t, err, "erro no serviço de CEP")
	assert.Empty(t, result)

	mockCep.AssertExpectations(t)
}

func TestExecute_ErroConsultaTemperatura(t *testing.T) {
	mockCep := new(mockCepService)
	mockTemp := new(mockTemperaturaService)

	cep := "88000000"
	cidade := "Florianópolis"

	mockCepResponse := consulta_cep.ConsultaCepResponse{
		Cep:    cep,
		Cidade: cidade,
	}

	mockCep.On("ConsultarCep", cep).Return(mockCepResponse, nil)
	mockTemp.On("ConsultaTemperatura", cidade).Return(consulta_temperatura.ConsultaTemperaturaResponse{}, errors.New("erro no serviço de temperatura"))

	useCase := NewConsultaTemperaturaUseCase(mockCep, mockTemp)

	result, err := useCase.Execute(cep)

	assert.Error(t, err)
	assert.EqualError(t, err, "erro no serviço de temperatura")
	assert.Empty(t, result)

	mockCep.AssertExpectations(t)
	mockTemp.AssertExpectations(t)
}
