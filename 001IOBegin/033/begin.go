// Известно, что X кг конфет стоит A рублей. Определить, сколько стоит
// 1 кг и Y кг этих же конфет
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y, z float64
	x = ioutil.NoNNumber("Введите X кол-во кг конфет")
	y = ioutil.NoNNumber("Введите их стоимость")
	fmt.Println("1 кг стоит:\t", y/x)
	z = ioutil.NoNNumber("Введите Y кол-во кг конфет")
	fmt.Printf("стоимость %v кг конфет:\t%v\n", z, z*(y/x))
}
