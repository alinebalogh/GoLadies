// passando argumentos dentro da função anônima
package exemplo_5

import "fmt"

func main() {

	func(num1 int) {
		fmt.Println(num1)
	}(4)

}
