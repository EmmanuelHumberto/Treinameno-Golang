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

	var nome = "Dono"
	var idade = 13

	//oprador:= usado para não declarar a palavra reservada var antes da variavel
	versao := 1.1

	fmt.Println("Olá sr!", nome, "Sua idade é:", idade, ". versão", versao)

	//reflect.TypeOf retorna o tipo do paramentro
	fmt.Println("O tipo da variavel nome é:", reflect.TypeOf(idade))
}
