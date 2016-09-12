// Дано целое число. Вывести его строку-описание вида «отрицательное чет-
// ное число», «нулевое число», «положительное нечетное число» и т. д.
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x int
	var s []string
	x = ioutil.Integer("целое число")

	if x == 0 {
		s = append(s, "нулевое ")
	} else {
		if x < 0 {
			s = append(s, "отрицательное ")
			if x%2 == 0 {
				s = append(s, "четное ")
			} else {
				s = append(s, "нечетное ")
			}
		} else {
			s = append(s, "положительное ")
			if x%2 == 0 {
				s = append(s, "четное ")
			} else {
				s = append(s, "нечетное ")
			}
		}
	}
	s = append(s, "число")
	fmt.Println(s)
}
