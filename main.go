package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {

	clienteDaniel := clientes.Titular{
		Nome:      "Daniel",
		CPF:       "12345312011",
		Profissao: "Programador",
	}

	contaDoDaniel := contas.ContaCorrente{
		Titular:       clienteDaniel,
		NumeroAgencia: 123,
		NumeroConta:   321,
	}

	contaDoDaniel.Depositar(432)
	PagarBoleto(&contaDoDaniel, 100)
	fmt.Println(contaDoDaniel)
}

func PagarBoleto(conta verificarConta, valor float64) {
	conta.Sacar(valor)
}

type verificarConta interface {
	Sacar(valor float64) string
}
