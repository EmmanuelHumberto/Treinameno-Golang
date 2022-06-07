//pacote principal
package main

//importando pacote:
import (
	// formataçao (impressão e formatação de saidas)
	"fmt"

	/*O pacote reflect implementa a reflexão em tempo de execução,
	permitindo que um programa manipule objetos com tipos arbitrários
	*/
	"reflect"
)

//função principal
func main() {

	//oprador:= usado para não declarar a palavra reservada var antes da variavel
	versao := 1.1

	//reflect.TypeOf retorna o tipo do paramentro
	fmt.Println("O tipo da variavel nome é:", reflect.TypeOf(versao))

	//monitoramentoewb
	fmt.Println("========================")
	fmt.Println("Menu de Opções ")
	fmt.Println("-------------------------")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("-------------------------")
	fmt.Println("2 - Exibir Log           ")
	fmt.Println("-------------------------")
	fmt.Println("3 - Sair do Programa     ")
	fmt.Println("-------------------------")

	comando := 0

	fmt.Println("Digite um numero para começar")
	// & =  endereço da variavel que o imput sera salvo
	fmt.Scanf("%d", &comando)

	if comando == 1 {
		fmt.Println("Monitoramento iniciado")

	} else if comando == 2 {
		fmt.Println("Carregando logs")

	} else if comando == 3 {
		fmt.Println("Saindo do Programa...")

	} else {
		fmt.Println("Comando digitado diferente das opções listadas")

	}

}
