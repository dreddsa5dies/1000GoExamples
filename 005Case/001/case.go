// Дано целое число в диапазоне 1–7. Вывести строку — название дня не-
// дели, соответствующее данному числу (1 — «понедельник», 2 — «втор-
// ник» и т. д.)
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x int
	for x > 7 || x < 1 {
		x = ioutil.Integer("число")
	}

	switch {
	case x == 1:
		fmt.Println("Понедельник")
	case x == 2:
		fmt.Println("Вторник")
	case x == 3:
		fmt.Println("Среда")
	case x == 4:
		fmt.Println("Четверг")
	case x == 5:
		fmt.Println("Пятница")
	case x == 6:
		fmt.Println("Суббота")
	case x == 7:
		fmt.Println("Воскресенье")
	}

}
