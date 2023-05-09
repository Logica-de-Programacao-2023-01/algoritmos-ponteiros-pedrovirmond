// Escreva uma função que receba um ponteiro para uma struct que contém informações
// de um produto (nome, preço e quantidade em estoque). A função deve atualizar o preço
// desse produto para um novo valor fornecido como argumento.
package main

import "fmt"

type Produto struct {
	Nome       string
	Preço      float64
	Quantidade int
}

func updatePreço(p *Produto, newPreço float64) {
	p.Preço = newPreço
}

func main() {
	p := Produto{"Calça jeans", 69.99, 43}
	fmt.Println("Preço antes da atualização: ", p.Preço)
	updatePreço(&p, 52.99)
	fmt.Println("Preço depois da atualização", p.Preço)
}
