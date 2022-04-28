package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		Abaixo temos todos os tipos que posso declarar uma variavel em Go
	*/
	var nome string = "Johnny"
	var versao = 1.1
	idade := 30
	fmt.Println("Olá sr.", nome, "sua idade é ", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("o tipo da variavel versao é ", reflect.TypeOf(versao))

	//como capturar no terminal uma entrada.
	var comando int
	fmt.Scanf("%d", &comando)
	//& -> aponta pro endereço de memoria dessa variavel
	fmt.Println("Endereço:", &comando)
	//foi criado uma função que não precisa que eu especifique que esta esperando um tipo determinado de dados
	//para isso foi criado a função Scan
	fmt.Scan(&comando)

}
