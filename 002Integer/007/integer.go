// Дано двузначное число. Найти сумму и произведение его цифр
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x uint
	for x < 10 || x > 100 {
		x = ioutil.UInteger("Введите двузначное число A")
	}
	fmt.Printf("сумма его цифр: %v\n", x/10+x%10)
	fmt.Printf("произведение его цифр: %v\n", (x/10)*(x%10))
}
