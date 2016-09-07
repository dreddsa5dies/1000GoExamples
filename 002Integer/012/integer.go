// Дано трехзначное число. Вывести число, полученное при прочтении
// исходного числа справа налево
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x uint
	for x < 100 || x > 1000 {
		x = util.UInteger("Введите трехзначное число A")
	}
	fmt.Printf("число справа налево: %v\n", (x%10)*100+10*((x-(100*(x/100)))/10)+x/100)
}
