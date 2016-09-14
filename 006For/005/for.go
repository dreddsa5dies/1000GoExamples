// Дано вещественное число — цена 1 кг конфет. Вывести стоимость 0.1,
// 0.2, ..., 1 кг конфет
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var m float64

	m = ioutil.Number("цена 1 кг конфет")
	fmt.Println("стоимость")
	for i := 0.1; i <= 1; i = i + 0.1 {
		fmt.Printf("%1.1f кг конфет\t %1.1f\n", i, m*i)
	}
}
