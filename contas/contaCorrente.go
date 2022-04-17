package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	podeSacar := valor > 0 && valor <= c.saldo

	if podeSacar {
		c.saldo -= valor
		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	podeDepositar := valor > 0

	if podeDepositar {
		c.saldo += valor
		return "Deposito realizado com sucesso.", c.saldo
	}

	return "Valor informado é incorreto", c.saldo
}

func (c *ContaCorrente) Transferir(valor float64, destino *ContaCorrente) string {
	podeTransferir := valor > 0 && valor <= c.saldo

	if podeTransferir {
		c.saldo -= valor
		destino.Depositar(valor)
		return "Valor transferido com sucesso"
	}

	return "Saldo insuficiente"
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
