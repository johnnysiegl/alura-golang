package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoJohnny := ContaCorrente{
		titular:       "Johnny",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.50,
	}

	contaDaAmanda := ContaCorrente{
		titular:       "Amanda",
		numeroAgencia: 333,
		numeroConta:   111333,
		saldo:         1020.50,
	}

	fmt.Println(contaDoJohnny)
	fmt.Println(contaDaAmanda)
}
