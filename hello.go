package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramento = 3
const delay = 5

func main() {
	exibirIntroducao()
	for {

		exibirMenu()

		comandoLido := lerComando()

		switch comandoLido {
		case 1:
			iniciarMonitoramento()
		case 2:
			print("Exibindo logs...")
		case 0:
			println("Saindo do programa...")
			os.Exit(0)
		default:
			println("Não conheço este comando")
			os.Exit(-1)

		}

	}

}

func exibirIntroducao() {
	nome := "Gustavo"
	versao := " 1.1"

	print("Olá, sr.", nome, "\n")
	print("Versão:", versao, "\n")
}
func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido) // Ponteiro
	println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func exibirMenu() {
	print("1 - Iniciar o monitoramento\n")
	print("2 - Exibir logs\n")
	print("0 - Sair\n")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://www.google.com.br", "https://www.alura.com.br", "https://www.caelum.com.br"}

	for I := 0; I < monitoramento; I++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testarSite(site)
		}
		time.Sleep(delay * time.Second)
		println("")
	}
	println("")
}

func testarSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Erro ao acessar o site:", err)
	}
	fmt.Println("Status do site:", resp.Status)
}
