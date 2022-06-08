//pacote principal
package main

//importando pacote:
import (
	//O pacote bufio implementa E/S em buffer.
	"bufio"
	//O pacote io fornece interfaces básicas para primitivas de E/S.
	//Seu trabalho principal é envolver as implementações existentes de tais primitivas,
	//como aquelas no pacote os
	"io"
	//O pacote ioutil implementa algumas funções de utilitário de E/S.
	"io/ioutil"
	//O pacote strconv implementa conversões de e para representações de string de tipos de dados básicos.
	"strconv"
	//As strings de pacote implementam funções simples para manipular strings codificadas em UTF-8.
	"strings"
	// formataçao (impressão e formatação de saidas)
	"fmt"
	//A rede de pacotes fornece uma interface portátil para E/S de rede, incluindo TCP/IP, UDP,
	//resolução de nomes de domínio e soquetes de domínio Unix.
	"net/http"
	//Pacote os responsavel pela conversa com os sistema opracional
	"os"
	//O tempo do pacote fornece funcionalidade para medir e exibir o tempo.
	//Os cálculos de calendário sempre assumem um calendário gregoriano, sem segundos bissextos.
	"time"
)

const daley = 5

//função principal
func main() {
	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			exibeLogs()
		case 3:
			sairDoPrograma()
		default:
			erro()
		}
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
		registraLog(site, true)
	} else {
		fmt.Println("Sites: problemas ao carregar. ", "Status Code",
			resp.StatusCode)
		registraLog(site, false)
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
func registraLog(site string, status bool) {
	//os.OpenFile( O_RDWR = abre o arquivo para leitura e escrita
	//O_CREATE cria o arquivo caso ele não exista. 0666 permissão para leitura e escrita
	//O_APPEND: anexa dados ao arquivo durante a escrita.
	arquivoLog, err := os.OpenFile("listaLog.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}
	//WriteString é como Write, mas grava o conteúdo da string s em vez de uma fatia de bytes.
	//strconv.FormatBool: converte de bool para string
	//time.Now().Format();Format retorna uma representação textual do valor de tempo
	//formatado de acordo com o layout definido pelo argumento
	arquivoLog.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - Online: " +
		strconv.FormatBool(status) + "\n")
	arquivoLog.Close()
}
func exibeLogs() {
	//O pacote ioutil implementa algumas funções de utilitário de E/S.
	//ReadFile lê o arquivo nomeado pelo nome do arquivo e retorna o conteúdo
	arquivoExibeLog, err := ioutil.ReadFile("listaLog.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	} else {
		fmt.Println(string(arquivoExibeLog))
	}
}
