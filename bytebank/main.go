package main

import (
	"/bytebank/contas"
	"fmt"
)

func main() {

	// contaDoVinicius := ContaCorrente{titular: "Vinicius", agencia: 1234, conta: 5678, saldo: 0}

	// contaDaRenata := ContaCorrente{"Renata", 1235, 5679, 0}
	// contaDaRenata2 := ContaCorrente{"Renata", 1235, 5679, 0}

	// fmt.Println(contaDoVinicius.titular, contaDaRenata, contaDaRenata2)

	// var contaDaCris *ContaCorrente
	// contaDaCris = new(ContaCorrente)
	// contaDaCris.titular = "Cris"
	// contaDaCris.saldo = 500

	// fmt.Println(*contaDaCris)

	contaDaSilvia := contas.ContaCorrente{titular: "Silvia", agencia: 1234, conta: 5678, saldo: 500}
	contaDoGustavo := contas.ContaCorrente{titular: "Gustavo", agencia: 1234, conta: 5679, saldo: 500}
	fmt.Println(contaDaSilvia.Sacar(100))
	contaDaSilvia.Depositar(500)
	fmt.Println(contaDaSilvia.saldo)
	fmt.Println(contaDoGustavo)
	fmt.Println(contaDaSilvia.Transferir(100, &contaDoGustavo))

}
