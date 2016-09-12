// Дано число A. Вычислить A^8 , используя вспомогательную перемен-
// ную и три операции умножения. Для этого последовательно находить A^2 ,
// A^4 , A^8 . Вывести все найденные степени числа A
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y float64
	x = ioutil.Number("Введите A")
	y = x * x
	fmt.Println("A^2:\t", y)
	y = y * y
	fmt.Println("A^4:\t", y)
	y = y * y
	fmt.Println("A^8:\t", y)
}
