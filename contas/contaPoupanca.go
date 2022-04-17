package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valor float64) string {
	podeSacar := valor > 0 && valor <= c.saldo

	if podeSacar {
		c.saldo -= valor
		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente"
}

func (c *ContaPoupanca) Depositar(valor float64) (string, float64) {
	podeDepositar := valor > 0

	if podeDepositar {
		c.saldo += valor
		return "Deposito realizado com sucesso.", c.saldo
	}

	return "Valor informado é incorreto", c.saldo
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
