// Дано трехзначное число. Вывести число, полученное при переста-
// новке цифр сотен и десятков исходного числа (например, 123 перейдет
// в 213)
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x uint
	for x < 100 || x > 1000 {
		x = ioutil.UInteger("Введите трехзначное число A")
	}
	fmt.Printf("полученное число: %v\n", 100*((x-(100*(x/100)))/10)+10*(x/100)+x%10)
}
