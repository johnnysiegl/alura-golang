package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 2

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			inicarMonitoramento()
		case 2:

		case 0:
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando.")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Johnny"
	versao := 1.0

	fmt.Println("Olá sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O valor da variável comando é:", comandoLido)

	return comandoLido
}

func inicarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := []string{"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br", "https://www.caelum.com.br"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
