// Дано число A. Вычислить A^15 , используя две вспомогательные пере-
// менные и пять операций умножения. Для этого последовательно находить
// A^2 , A^3 , A^5 , A^10 , A^15 . Вывести все найденные степени числа A
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y, z float64
	x = util.Number("Введите A")
	y = x * x
	fmt.Println("A^2:\t", y)
	z = y * x
	fmt.Println("A^3:\t", z)
	y = z * y
	fmt.Println("A^5:\t", y)
	x = y * y
	fmt.Println("A^10:\t", x)
	z = x * y
	fmt.Println("A^15:\t", z)
}
