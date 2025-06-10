package outbound

import (
	"errors"
	"fmt"
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
	var (
		maxAttempts = 5
		baseDelay   = 100 * time.Millisecond
		maxDelay    = 2 * time.Second
		resp        PrecoResponse
		err         error
		delay       = baseDelay
	)

	for i := 1; i <= maxAttempts; i++ {
		// Simula delay entre 100ms e 2s
		simulatedDelay := time.Duration(100+rand.Intn(1900)) * time.Millisecond
		time.Sleep(simulatedDelay)

		// Simula falha aleatória (~30% das vezes)
		if rand.Float64() < 0.3 {
			err = errors.New("erro ao consultar fornecedor dummy")
			fmt.Printf("[Retry] Tentativa %d: Falhou (%v). Vai tentar novamente em %v\n", i, err, delay)
		} else {
			//Se der sucesso, retorna um preço fake
			preco := 10.0 + rand.Float64()*90.0 //Preço entre 10 e 100
			resp = PrecoResponse{
				Produto:    produto,
				Preco:      preco,
				Moeda:      "BRL",
				Fornecedor: f.Nome,
			}
			return resp, nil
		}

		if i < maxAttempts {
			//Backoff exponencial + jitter
			jitter := time.Duration(rand.Int63n(int64(delay)))
			sleep := delay + jitter
			if sleep > maxDelay {
				sleep = maxDelay
			}
			time.Sleep(sleep)
			delay *= 2
			if delay > maxDelay {
				delay = maxDelay
			}
		}
	}
	// Todas tentativas falharam
	return PrecoResponse{}, fmt.Errorf("todas as tentativas falharam: %w", err)
}
