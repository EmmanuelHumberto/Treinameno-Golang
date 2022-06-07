//pacote principal
package main

//importando pacote:
import (
	// formataçao (impressão e formatação de saidas)
	"fmt"
	"net/http"
	"os"
	"reflect"
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
	versao := 1.1
	//reflect.TypeOf retorna o tipo do paramentro
	fmt.Println("versão", versao)
	fmt.Println("O tipo da variavel nome é:", reflect.TypeOf(versao))
}
func leComando() int {
	var comandoLido int
	fmt.Println("Digite um numero para começar")
	// & =  endereço da variavel que o imput sera salvo
	fmt.Scanf("%d", &comandoLido)
	return comandoLido

}
func exibeMenu() {
	//menu
	fmt.Println("========================")
	fmt.Println("Menu de Opções ")
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

	for {
		fmt.Println("Monitoramento iniciado")
		site := "https://random-status-code.herokuapp.com/"
		http.Get(site)
		//operador _ inativa uma das variaveis de multiplos retornos
		resp, _ := http.Get(site)

		if resp.StatusCode == 200 {
			fmt.Println("Sites:", site, " carregados com sucesso")
		} else {
			fmt.Println("Sites:", site, " problemas ao carregar. ", "Status Code",
				resp.StatusCode)
		}
	}
}
