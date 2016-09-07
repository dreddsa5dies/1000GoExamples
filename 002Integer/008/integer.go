// Дано двузначное число. Вывести число, полученное при перестанов-
// ке цифр исходного числа
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x uint
	for x < 10 || x > 100 {
		x = util.UInteger("Введите двузначное число A")
	}
	fmt.Printf("перестановка цифр исходного числа: %v\n", 10*(x%10)+(x/10))
}
