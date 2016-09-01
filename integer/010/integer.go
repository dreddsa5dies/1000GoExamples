// ДДано трехзначное число. Вывести вначале его последнюю цифру
// (единицы), а затем — его среднюю цифру (десятки)
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
	fmt.Printf("его последняя цифра (единицы): %v\n", x%10)
	fmt.Printf("его средняя цифра (десятки): %v\n", (x-(100*(x/100)))/10)
}
