// Описать процедуру DigitCountSum( K , C , S ), находящую количество C
// цифр целого положительного числа K , а также их сумму S ( K — входной,
// C и S — выходные параметры целого типа). С помощью этой процедуры
// найти количество и сумму цифр для каждого из пяти данных целых чисел

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var k uint
	for i := 0; i < 5; i++ {
		k = ioutil.UInteger("целое положительное число")
		c, s := digitCountSum(k)
		fmt.Printf("количество цифр числа\t: %v\n", c)
		fmt.Printf("сумма этих цифр\t: %v\n", s)
	}
}

func digitCountSum(k uint) (c, s uint) {
	c = 0
	s = 0
	for k != 0 {
		s += k % 10
		c++
		k = k / 10
	}
	return c, s
}
