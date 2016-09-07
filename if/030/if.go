// Дано целое число, лежащее в диапазоне 1–999. Вывести его строку-
// описание вида «четное двузначное число», «нечетное трехзначное число»
// и т. д.
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x uint
	var s []string
	for x < 1 || x > 999 {
		x = util.UInteger("целое число, лежащее в диапазоне 1–999")
	}

	if x%2 == 0 {
		s = append(s, "четное")
	} else {
		s = append(s, "нечетное")
	}

	if x < 10 {
		s = append(s, "однозначное")
	} else {
		if x > 9 && x < 100 {
			s = append(s, "двухзначное")
		} else if x > 99 && x < 1000 {
			s = append(s, "трехзначное")
		}
	}

	s = append(s, "число")
	fmt.Println(s)
}
