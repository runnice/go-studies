package main

import (
	"bytebank/contas"
	"fmt"
)

func main() {

	// contaDoVinicius := ContaCorrente{Titular: "Vinicius", agencia: 1234, conta: 5678, Saldo: 0}

	// contaDaRenata := ContaCorrente{"Renata", 1235, 5679, 0}
	// contaDaRenata2 := ContaCorrente{"Renata", 1235, 5679, 0}

	// fmt.Println(contaDoVinicius.Titular, contaDaRenata, contaDaRenata2)

	// var contaDaCris *ContaCorrente
	// contaDaCris = new(ContaCorrente)
	// contaDaCris.Titular = "Cris"
	// contaDaCris.Saldo = 500

	// fmt.Println(*contaDaCris)

	contaDaSilvia := contas.ContaCorrente{Titular: "Joaquim", Agencia: 1234, Conta: 5678, Saldo: 0}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Agencia: 1234, Conta: 5679, Saldo: 500}
	fmt.Println(contaDaSilvia.Sacar(100))
	contaDaSilvia.Depositar(500)
	fmt.Println(contaDaSilvia.Saldo)
	fmt.Println(contaDoGustavo)
	fmt.Println(contaDaSilvia.Transferir(100, &contaDoGustavo))

}
