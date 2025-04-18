package consulta_temperatura

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

func TestWeatherApi_ConsultaTemperatura_Success(t *testing.T) {
	mockRequest := new(MockHttpRequest)
	mockRequest.On("IsValid").Return(nil)
	mockRequest.On("Execute").Return(map[string]interface{}{
		"location": map[string]interface{}{
			"name": "São Paulo",
		},
		"current": map[string]interface{}{
			"temp_c": 25.0,
		},
	}, nil)

	requestFactory := func(method, url string) external_http_request.ExternalHttpRequestInterface {
		return mockRequest
	}

	weatherApi := NewWeatherApi("test-api-key", requestFactory)
	response, err := weatherApi.ConsultaTemperatura("São Paulo")

	assert.NoError(t, err)
	assert.Equal(t, "São Paulo", response.Location.Name)
	assert.Equal(t, 25.0, response.Current.TempC)

	mockRequest.AssertExpectations(t)
}

func TestWeatherApi_ConsultaTemperatura_InvalidRequest(t *testing.T) {
	mockRequest := new(MockHttpRequest)
	mockRequest.On("IsValid").Return(errors.New("invalid request"))

	requestFactory := func(method, url string) external_http_request.ExternalHttpRequestInterface {
		return mockRequest
	}

	weatherApi := NewWeatherApi("test-api-key", requestFactory)
	_, err := weatherApi.ConsultaTemperatura("São Paulo")

	assert.Error(t, err)
	assert.Equal(t, "invalid request", err.Error())

	mockRequest.AssertExpectations(t)
}

func TestWeatherApi_ConsultaTemperatura_ExecuteError(t *testing.T) {
	mockRequest := new(MockHttpRequest)
	mockRequest.On("IsValid").Return(nil)
	mockRequest.On("Execute").Return(nil, errors.New("execution error"))

	requestFactory := func(method, url string) external_http_request.ExternalHttpRequestInterface {
		return mockRequest
	}

	weatherApi := NewWeatherApi("test-api-key", requestFactory)
	_, err := weatherApi.ConsultaTemperatura("São Paulo")

	assert.Error(t, err)
	assert.Equal(t, "execution error", err.Error())

	mockRequest.AssertExpectations(t)
}
