// Escreva uma função que receba um ponteiro para uma struct que contém informações de um produto (nome, preço e quantidade em estoque). A função deve atualizar o preço desse produto para um novo valor fornecido como argumento.

package main

import "fmt"

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

func refreshPrice(p *Produto, newPrice float64) {
	p.Preco = newPrice
}

func main() {
	produto := Produto{
		Nome:       "Bakugan",
		Preco:      10,
		Quantidade: 30,
	}
	fmt.Println("Preço antigo:", produto.Preco)

	refreshPrice(&produto, 15)
	fmt.Print("Preço novo: ", produto.Preco)

}
