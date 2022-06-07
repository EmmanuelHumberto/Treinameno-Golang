//pacote principal
package main

//importando pacote:
import (
	"bufio"
	"io"
	"strings"

	// formataçao (impressão e formatação de saidas)
	"fmt"
	"net/http"
	"os"
	"time"
)

const daley = 60

//função principal
func main() {
	exibeIntroducao()

	exibeMenu()
	comando := leComando()
	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Carregando logs")
	case 3:
		sairDoPrograma()
	default:
		erro()
	}
}
func exibeIntroducao() {
	//oprador:= usado para não declarar a palavra reservada var antes da variavel
	versao := float32(1.1)
	fmt.Println("========================")
	//reflect.TypeOf retorna o tipo do paramentro
	fmt.Println("versão", versao)
}
func leComando() int {
	var comandoLido int
	fmt.Println("")
	fmt.Println("Digite um numero para começar")
	// & =  endereço da variavel que o imput sera salvo
	fmt.Scanf("%d", &comandoLido)
	return comandoLido
}
func exibeMenu() {
	//menu
	fmt.Println("========================")
	fmt.Println(" *******  MENU  ******* ")
	fmt.Println("-------------------------")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("-------------------------")
	fmt.Println("2 - Exibir Log           ")
	fmt.Println("-------------------------")
	fmt.Println("3 - Sair do Programa     ")
	fmt.Println("========================")
}
func sairDoPrograma() {
	fmt.Println("Programa Finalizado com Sucesso!")
	os.Exit(0)
}
func erro() {
	fmt.Println("Comando digitado diferente das opções listadas")
	os.Exit(-1)
}
func iniciarMonitoramento() {
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("Monitoramento iniciado")
	fmt.Println("-------------------------------------------------------------")
	sites := leSitesArquivo()

	monitoramntos := len(sites)
	for c := 1; c <= monitoramntos; c++ {
		for i, site := range sites {
			fmt.Println("Testando site:", i, " : ", site)
			testaSite(site)
			fmt.Println("-------------------------------------------------------------")
		}
		fmt.Println("Contagem: ", c)
		fmt.Println("_____________________________________________________________")
		time.Sleep(daley * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}
func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Sites: carregados com sucesso")
	} else {
		fmt.Println("Sites: problemas ao carregar. ", "Status Code",
			resp.StatusCode)
	}
}
func leSitesArquivo() []string {
	var sites []string
	//Abrindo arquivo com so.Open
	arquivoSite, err := os.Open("listaSites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}
	//bufio -> O pacote bufio implementa E/S em buffer.
	//bufio.NewReader retorna um leitor que irá ler linha a linha do arquivo
	leitor := bufio.NewReader(arquivoSite)
	for {
		//retorna a linha lida em uma string1
		linha, err := leitor.ReadString('\n')
		//TriSpece retira so epaços vazios e os \n ao final de cada linha
		linha = strings.TrimSpace(linha)
		//A função interna append anexa elementos ao final de uma fatia
		sites = append(sites, linha)
		fmt.Println(linha)
		//IOF erro de final de arquivo
		if err == io.EOF {
			break
		}
	}
	//Fechando arquivo apos a leitura
	arquivoSite.Close()
	return sites
}
