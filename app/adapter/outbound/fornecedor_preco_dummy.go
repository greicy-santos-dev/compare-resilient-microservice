package outbound

import (
	"errors"
	"math/rand"
	"time"
)

type FornecedorPrecoDummy struct {
	Nome string
}

func NewFornecedorPrecoDummy(nome string) *FornecedorPrecoDummy {
	return &FornecedorPrecoDummy{
		Nome: nome,
	}
}

func (f *FornecedorPrecoDummy) GetPreco(produto string) (PrecoResponse, error) {
	// Simula delay entre 100ms e 2s
	delay := time.Duration(100+rand.Intn(1900)) * time.Millisecond
	time.Sleep(delay)

	// Simula falha aleatória (~30% das vezes)
	if rand.Float64() < 0.3 {
		return PrecoResponse{}, errors.New("erro ao consultar fornecedor dummy")
	}

	//Se der sucesso, retorna um preço fake
	preco := 10.0 + rand.Float64()*90.0 //Preço entre 10 e 100
	return PrecoResponse{
		Produto:    produto,
		Preco:      preco,
		Moeda:      "BRL",
		Fornecedor: f.Nome,
	}, nil
}
