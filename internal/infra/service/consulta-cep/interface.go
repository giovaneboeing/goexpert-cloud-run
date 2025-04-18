package consulta_cep

type ConsultaCepResponse struct {
	Cep         string `json:"cep,omitempty"`
	Logradouro  string `json:"logradouro,omitempty"`
	Complemento string `json:"complemento,omitempty"`
	Unidade     string `json:"unidade,omitempty"`
	Bairro      string `json:"bairro,omitempty"`
	Cidade      string `json:"localidade,omitempty"`
	UF          string `json:"uf,omitempty"`
	Estado      string `json:"estado,omitempty"`
	Regiao      string `json:"regiao,omitempty"`
	Ibge        string `json:"ibge,omitempty"`
	Gia         string `json:"gia,omitempty"`
	DDD         string `json:"ddd,omitempty"`
	Siafi       string `json:"siafi,omitempty"`
	Erro        string `json:"erro,omitempty"`
}

type CepServiceInterface interface {
	ConsultarCep(cep string) (ConsultaCepResponse, error)
}
