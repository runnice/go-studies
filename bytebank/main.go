package main

import "fmt"

type ContaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

func main() {
	contaDoVinicius := ContaCorrente{titular: "Vinicius", agencia: 1234, conta: 5678, saldo: 0}

	contaDaRenata := ContaCorrente{"Renata", 1235, 5679, 0}
	contaDaRenata2 := ContaCorrente{"Renata", 1235, 5679, 0}

	fmt.Println(contaDoVinicius.titular, contaDaRenata, contaDaRenata2)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	fmt.Println(*contaDaCris)

}
