package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	nome := "Gustavo"
	var versao float32 = 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Versão:", versao)

	fmt.Println("1 - Iniciar o monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	if comando == 1 {
		iniciarMonitoramento()
	}
	if comando == 2 {
		imprimirLogs()
		os.Exit(0)

	}
	if comando == 0 {
		fmt.Println("Saindo...")
		os.Exit(0)
	} else {
		fmt.Println("Comando não reconhecido")
		os.Exit(-1)
	}
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := lerSitesArquivos()

	for i, site := range sites {
		fmt.Println("Testando site", i, ":", site)
		testarSite(site)
	}
}

func lerSitesArquivos() []string {
	var sites []string
	arquivo, err := os.ReadFile("sites.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return sites
	}

	linhas := strings.Split(string(arquivo), "\n")
	for _, linha := range linhas {
		linha = strings.TrimSpace(linha)
		if linha != "" {
			if !strings.HasPrefix(linha, "http://") && !strings.HasPrefix(linha, "https://") {
				linha = "https://" + linha
			}
			sites = append(sites, linha)
		}
	}

	return sites
}

func testarSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Erro ao acessar o site:", err)
		return
	}
	if resp.StatusCode == 200 {
		println("Site", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		println("Site", site, "está com problemas. Status:", resp.Status)
		registraLog(site, false)
	}
	fmt.Println("Status do site:", resp.Status)
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de log:", err)
		return
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 ") + site + " - online: " + fmt.Sprintf("%t", status) + "\n")

	arquivo.Close()

}
func imprimirLogs() {
	arquivo, err := os.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de log:", err)
		return
	}

	fmt.Println(string(arquivo))

}
