// Exemplo de uma função retornando vários valores
package exemplo_1

import "fmt"

func myfunc(p, q int) (int, int, int) {
	return p - q, p * q, p + q
}

func main() {

	var myvar1, myvar2, myvar3 = myfunc(4, 2)

	fmt.Println("Result is:", myvar1)
	fmt.Println("Result is:", myvar2)
	fmt.Println("Result is:", myvar3)
}
