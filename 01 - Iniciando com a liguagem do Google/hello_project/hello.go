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

	//03
	//no go eu não preciso colocar parenteses
	//toda checagem de (if) deve retornar um true ou false (todas expressão)
	// if comando == 1 {

	// } else if comando == 2 {

	// } else if comando == 0 {

	// }
	//a linhagem go não tem o break para controlar o switch
	// switch comando {
	// case 1:

	// case 2:

	// case 0:

	// default:

	// }

	//funções
	//Elas podem retornar mais de um valor. é comum que isso aconteça, porque normalmente elas possuem um retorno de erro.
	// resp, err := http.Get(site)
	//Para ignorar um dos retornos de alguma função, você pode apenas utilizar o "_" para dizer para o Go ignorar aquele retorno
	//  resp, _ := http.Get(site)

	//coleções
	//array em go sempre vai ser tamanho fixo
	//slice ele vai ser dinamico, mas por baixo dos panos ele trabalha com array
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br"

	exibeNomes()
	//o slice, quando você adiciona um novo item no slice que vc ja havia criado, ele dobra de tamanho
}

func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Bernardo"}
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	nomes = append(nomes, "Aparecida")
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
}
