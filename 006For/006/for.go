// Дано вещественное число — цена 1 кг конфет. Вывести стоимость 1.2,
// 1.4, ..., 2 кг конфет.
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var m float64

	m = ioutil.Number("цена 1 кг конфет")
	fmt.Println("стоимость")
	for i := 1.2; i <= 2; i = i + 0.2 {
		fmt.Printf("%1.1f кг конфет\t %1.1f\n", i, m*i)
	}
}
