package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/greicy-santos-dev/compare-resilient-microservice/app/adapter/outbound"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // importante para random funcionar corretamente

	fornecedor := outbound.NewFornecedorPrecoDummy("Fornecedor Dummy")

	produto := "Arroz"

	preco, err := fornecedor.GetPreco(produto)
	if err != nil {
		fmt.Printf("Erro ao consultar preço: %v\n", err)
	} else {
		fmt.Printf("Preço do produto %s no %s: R$ %.2f %s\n", preco.Produto, preco.Fornecedor, preco.Preco, preco.Moeda)
	}
}
