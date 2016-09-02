// Даны три целых числа: A , B , C . Проверить истинность высказыва-
// ния: «Каждое из чисел A , B , C положительное»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var a, b int
	a = util.Integer("целое число A")
	b = util.Integer("целое число B")
	fmt.Println((a%2 != 0 && b%2 == 0) || (a%2 == 0 && b%2 != 0))
}
