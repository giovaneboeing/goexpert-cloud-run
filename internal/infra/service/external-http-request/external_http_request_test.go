package external_http_request

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIsValid_Success(t *testing.T) {
	req := NewHttpRequest("GET", "http://example.com")
	err := req.IsValid()
	assert.NoError(t, err)
}

func TestIsValid_InvalidMethod(t *testing.T) {
	req := NewHttpRequest("", "http://example.com")
	err := req.IsValid()
	assert.EqualError(t, err, "invalid method")
}

func TestIsValid_InvalidUrl(t *testing.T) {
	req := NewHttpRequest("GET", "")
	err := req.IsValid()
	assert.EqualError(t, err, "invalid url")
}

func TestExecute_Success(t *testing.T) {
	expected := map[string]string{"message": "ok"}

	// cria servidor de teste
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(expected)
	}))
	defer ts.Close()

	req := NewHttpRequest("GET", ts.URL)
	result, err := req.Execute()

	assert.NoError(t, err)

	// o retorno é interface{}, então convertemos pra map
	resultMap, ok := result.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, expected["message"], resultMap["message"])
}

func TestExecute_InvalidRequest(t *testing.T) {
	req := NewHttpRequest("", "") // inválido
	_, err := req.Execute()
	assert.EqualError(t, err, "invalid method")
}

func TestExecute_HttpError(t *testing.T) {
	// endereço inválido para forçar erro de conexão
	req := NewHttpRequest("GET", "http://127.0.0.1:0")
	_, err := req.Execute()
	assert.Error(t, err)
}

func TestExecute_StatusNotOK(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	}))
	defer ts.Close()

	req := NewHttpRequest("GET", ts.URL)
	_, err := req.Execute()
	assert.EqualError(t, err, "failed to execute request")
}

func TestExecute_InvalidJsonResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("invalid json"))
	}))
	defer ts.Close()

	req := NewHttpRequest("GET", ts.URL)
	_, err := req.Execute()
	assert.Error(t, err)
}
