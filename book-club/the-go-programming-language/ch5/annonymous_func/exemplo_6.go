// passando uma função anônima como argumento para outra função
package exemplo_6

import "fmt"

// está sendo passado como argumento a função anônima "func(p, q)"
func GFL(i func(p, q string) string) {
	fmt.Println(i("Go", "for"))

}

func main() {
	value := func(p, q string) string {
		return p + q + "Ladies"
	}
	GFL(value)
}
