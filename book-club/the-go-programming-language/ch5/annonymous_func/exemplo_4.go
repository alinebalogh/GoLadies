// usando função anônima atribuindo a uma variável
package exemplo_4

import "fmt"

func main() {

	// criando uma variável chamada value e chamando ela dentro da função main
	value := func() {
		fmt.Println("Welcome! to Goladies")
	}
	value()

}
