// Дано целое число K . Вывести строку-описание оценки, соответствую-
// щей числу K (1 — «плохо», 2 — «неудовлетворительно», 3 — «удовлетво-
// рительно», 4 — «хорошо», 5 — «отлично»). Если K не лежит в диапазоне
// 1–5, то вывести строку «ошибка»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x int
	x = util.Integer("число")

	switch {
	case x == 1:
		fmt.Println("плохо")
	case x == 2:
		fmt.Println("неудовлетворительно")
	case x == 3:
		fmt.Println("удовлетворительно")
	case x == 4:
		fmt.Println("Хорошо")
	case x == 5:
		fmt.Println("Отлично")
	default:
		fmt.Println("Ошибка")
	}
}
