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

	//criando variaveis nulas
	var s *string = nil
	fmt.Println(s)

	/*
		Ponteiros:
	*/
	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"

	contaDoJohnny2 := ContaCorrente{
		titular:       "Johnny",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.50,
	}
	fmt.Println(contaDoJohnny == contaDoJohnny2)

	contaDoJohnny2.numeroConta = 899
	fmt.Println(contaDoJohnny == contaDoJohnny2)

	var contaDaCris2 *ContaCorrente
	contaDaCris2 = new(ContaCorrente)
	contaDaCris2.titular = "Cris"
	fmt.Println(contaDaCris)
	fmt.Println(contaDaCris2)
	fmt.Println(&contaDaCris)
	fmt.Println(&contaDaCris2)
	fmt.Println(contaDaCris == contaDaCris2)
	fmt.Println(*contaDaCris == *contaDaCris2)

	//
	fmt.Println(Somando(1))
	fmt.Println(Somando(1, 1))
	fmt.Println(Somando(1, 1, 1))
	fmt.Println(Somando(1, 1, 2, 4))
}

//função com parametros inderterminados
func Somando(numeros ...int) int {
	resultadoDaSoma := 0
	for _, numero := range numeros {
		resultadoDaSoma += numero
	}
	return resultadoDaSoma
}
