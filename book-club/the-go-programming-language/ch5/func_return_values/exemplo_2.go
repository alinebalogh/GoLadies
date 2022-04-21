package exemplo_2

import "fmt"

// Poderia ser "func myfunc(p, q int) (rectangle, square int)"
func myfunc(p, q int) (rectangle int, square int) {
	rectangle = p * q
	square = p * p
	return
}

func main() {
	var area1, area2 = myfunc(2, 4)

	fmt.Println("Area of the rectangle is:", area1)
	fmt.Println("Area of the square is:", area2)
}
