package consulta_cep

import (
	"errors"
	external_http_request "github.com/giovaneboeing/desafio-cloud-run/internal/infra/service/external-http-request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockHttpRequest struct {
	mock.Mock
}

func (m *MockHttpRequest) IsValid() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockHttpRequest) Execute() (interface{}, error) {
	args := m.Called()
	return args.Get(0), args.Error(1)
}

func TestViaCep_ConsultarCep_Success(t *testing.T) {
	mockRequest := new(MockHttpRequest)
	mockRequest.On("IsValid").Return(nil)
	mockRequest.On("Execute").Return(map[string]interface{}{
		"cep":        "01001-000",
		"logradouro": "Praça da Sé",
		"bairro":     "Sé",
		"localidade": "São Paulo",
		"uf":         "SP",
	}, nil)

	requestFactory := func(method, url string) external_http_request.ExternalHttpRequestInterface {
		return mockRequest
	}

	viaCep := NewViaCep(requestFactory)
	response, err := viaCep.ConsultarCep("01001000")

	assert.NoError(t, err)
	assert.Equal(t, "01001-000", response.Cep)
	assert.Equal(t, "Praça da Sé", response.Logradouro)
	assert.Equal(t, "Sé", response.Bairro)
	assert.Equal(t, "São Paulo", response.Cidade)
	assert.Equal(t, "SP", response.UF)

	mockRequest.AssertExpectations(t)
}

func TestViaCep_ConsultarCep_InvalidRequest(t *testing.T) {
	mockRequest := new(MockHttpRequest)
	mockRequest.On("IsValid").Return(errors.New("invalid request"))

	requestFactory := func(method, url string) external_http_request.ExternalHttpRequestInterface {
		return mockRequest
	}

	viaCep := NewViaCep(requestFactory)
	_, err := viaCep.ConsultarCep("01001000")

	assert.Error(t, err)
	assert.Equal(t, "invalid request", err.Error())

	mockRequest.AssertExpectations(t)
}

func TestViaCep_ConsultarCep_ExecuteError(t *testing.T) {
	mockRequest := new(MockHttpRequest)
	mockRequest.On("IsValid").Return(nil)
	mockRequest.On("Execute").Return(nil, errors.New("execution error"))

	requestFactory := func(method, url string) external_http_request.ExternalHttpRequestInterface {
		return mockRequest
	}

	viaCep := NewViaCep(requestFactory)
	_, err := viaCep.ConsultarCep("01001000")

	assert.Error(t, err)
	assert.Equal(t, "execution error", err.Error())

	mockRequest.AssertExpectations(t)
}
