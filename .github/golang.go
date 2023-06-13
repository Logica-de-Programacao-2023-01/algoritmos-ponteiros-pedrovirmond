//Escreva uma função em Go que receba um ponteiro
//para um struct Conta com campos saldo e titular,
//e adicione um valor ao saldo da conta.
package main

import "fmt"

type Conta struct {
	Saldo   float64
	Titular string
}

func addValor(conta *Conta) {
	var valor float64
	novoSaldo := conta.Saldo + valor
}

func main() {
	conta := Conta{
		Saldo:   4120.21,
		Titular: "Pedro",
	}

	var valor float64
	fmt.Print("Apresente o valor qu deseja adicionar: ")
	fmt.Scanln(&valor)

	addValor(&conta)
	fmt.Println("")

}