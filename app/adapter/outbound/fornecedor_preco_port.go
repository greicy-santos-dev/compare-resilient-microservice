package outbound

type PrecoResponse struct {
	Produto    string  `json:"produto"`
	Preco      float64 `json:"preco"`
	Moeda      string  `json:"moeda"`
	Fornecedor string  `json:"fornecedor"`
}

type FornecedorPrecoPort interface {
	GetPreco(produto string) (*PrecoResponse, error)
}
