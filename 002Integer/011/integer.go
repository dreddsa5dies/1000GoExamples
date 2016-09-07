// Дано трехзначное число. Найти сумму и произведение его цифр
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
	fmt.Printf("сумма его цифр: %v\n", x/100+(x-(100*(x/100)))/10+x%10)
	fmt.Printf("произведение его цифр: %v\n", (x/100)*((x-(100*(x/100)))/10)*(x%10))
}
