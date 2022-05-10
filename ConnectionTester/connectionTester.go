package main

import (
	"bufio"
	"fmt" // pacote printa na tela e formata
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	exibeIntroducao()
	leSitesdoArquivo()
	registraLog("site-falso", false)
	for {
		exibeMenu()

		// var nome string = "Runas" //declarando variavel
		// var idade int = 34
		// var versao float32 = 1.1

		// := declara variável e infere tipo automaticamente

		//Menu com If Else

		// if comando == 1 {
		// 	fmt.Println("Monitorando...")
		// } else if comando == 2 {
		// 	fmt.Println("Exibindo logs...")
		// } else if comando == 0 {
		// 	fmt.Println("Saindo do programa...")
		// } else {
		// 	fmt.Println("Não conheço este comando")
		// }

		switch lerComando() {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0) // Encerrar o programa
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

	// fmt.Println("Endereço da variável comando", &comando)
	// fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
	// fmt.Println("O tipo da variável versão é", reflect.TypeOf(versao))
}

func exibeIntroducao() {
	nome := "Runas" //declarando variavel
	versao := 1.1

	fmt.Println("Hello Mr.", nome)                      //imprimir na tela
	fmt.Println("Este programa está na versão", versao) //imprimir na tela
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir os logs")
	fmt.Println("0- Sair do programa")

}
func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido) //Primeiro sempre o modificador, segundo endereço (ponteiro) na variável & = aponta para o ponteiro da variável
	fmt.Println("O comando escolhido foi", comandoLido)

	//Diferenças entre Scanf e Scan
	//Scan já espera o int então só aceitará Int
	// fmt.Scanf("%d", &comando) //Primeiro sempre o modificador, segundo endereço (ponteiro) na variável & = aponta para o ponteiro da variável

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://random-status-code.herokuapp.com", "https://www.alura.com.br", "https://www.caelum.com.br"}
	// sites[0] = "https://random-status-code.herokuapp.com"
	// sites[1] = "https://www.alura.com.br"
	// sites[2] = "https://www.caelum.com.br"

	fmt.Println(sites)
	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites { //range vai verificar posição(i) e quem está naquela posição(site)
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")

}
func testaSite(site string) {
	resp, err := http.Get(site)
	// fmt.Println(resp)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site", site, "está com problemas. StatusCode:", resp.StatusCode)
		registraLog(site, false)
	}

}

func leSitesdoArquivo() []string {
	var sites []string

	// arquivo, err := ioutil.ReadFile("sites.txt") //imprimir todo arquivo
	// fmt.Println(string(arquivo)) //usando ioutil

	arquivo, err := os.Open("sites.txt") //abre o arquivo
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)

	}
	fmt.Println("Arquivo aberto com sucesso!")
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}

	}

	fmt.Println(sites)
	fmt.Println(len(sites))
	arquivo.Close() //fecha o arquivo

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

//Uso de Slice
// func exibeNomes(){
// 	nomes:= []string{"Douglas", "Daniel", "Bernardo"}
//  nomes = append(nomes, "Aparencida") //vai dar push no nome Aparecida para o Slice nomes
// 	fmt.Println(nomes)
// 	fmt.Println("o meu slice tem", len(nomes))
// 	fmt.Println("o meu slice tem capacidade de", cap(nomes))
// }

// Fazer função retornar 2 valores

//nome, idade := devolveNomeEIdade()
//_, idade := devolveNomeEIdade() // Underline ignora a obrigatoriedade de receber o primeira variável PS: Funciona em outras tmb. Pode ser na segunda ou terceira variável.
//.Println(nome, idade)

// func devolveNomeEIdade() (string, int) {
// 	nome := "Runas"
// 	idade := 34
// 	return nome, idade
//}
