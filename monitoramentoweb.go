//pacote principal
package main

//importando pacote:
import (
	// formataçao (impressão e formatação de saidas)
	"fmt"
	"net/http"
	"os"
	"time"
)

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
	fmt.Println("MENU DE OPÇOES ")
	fmt.Println("-------------------------")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("-------------------------")
	fmt.Println("2 - Exibir Log           ")
	fmt.Println("-------------------------")
	fmt.Println("3 - Sair do Programa     ")
	fmt.Println("-------------------------")
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
	fmt.Println("Monitoramento iniciado")
	sites := []string{
		"https://random-status-code.herokuapp.com/",
		"https://www.vallum.com.br",
		"https://www.google.com.br"}
	monitoramntos := len(sites)
	for c := 1; c < monitoramntos; c++ {
		for i, site := range sites {
			fmt.Println("Testando site:", i, " : ", site)
			testaSite(site)
		}
		fmt.Println("Contagem: ", c)
		fmt.Println("-------------------------")
		time.Sleep(5 * time.Second)
	}
	fmt.Println("")
}
func testaSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Sites: carregados com sucesso")
	} else {
		fmt.Println("Sites: problemas ao carregar. ", "Status Code",
			resp.StatusCode)
	}
}
