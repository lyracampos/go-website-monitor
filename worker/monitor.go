package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 5
const delay = 3

func main() {
	exibirIntroducao()

	for {
		exibirMenu()
		comando := lerComando()
		fmt.Println("")

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			exibirLogs()
		case 0:
			fmt.Println("Sair do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando desconhecido")
			os.Exit(-1)
		}
	}
}

func exibirIntroducao() {
	nome := "Guilherme"
	versao := 1.0
	fmt.Println("Olá, sr.", nome)
	fmt.Println("O programa está na versão:", versao)
}

func exibirMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func lerComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func iniciarMonitoramento() {

	sites := lerSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		fmt.Println("iniciando teste: ", i+1)
		for i, site := range sites {
			fmt.Println("testando site", i, ":", site)
			testarSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testarSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("ocorreu um erro: ", err)
	} else {
		if resp.StatusCode == 200 {
			fmt.Println("o site", site, "está funcionando com sucesso")
			registrarLog(site, true)
		} else {
			fmt.Println("o site", site, "apresentou o seguinte erro:", resp.StatusCode)
			registrarLog(site, false)
		}
	}
}

func exibirLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(string(arquivo))
}

func lerSitesDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registrarLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("ocorreu um erro: ", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}
