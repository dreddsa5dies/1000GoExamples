// Единицы массы пронумерованы следующим образом: 1 — килограмм,
// 2 — миллиграмм, 3 — грамм, 4 — тонна, 5 — центнер. Дан номер единицы
// массы (целое число в диапазоне 1–5) и масса тела в этих единицах (веще-
// ственное число). Найти массу тела в килограммах
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var (
		x uint
		y float64
	)

	for x < 1 || x > 5 {
		x = util.UInteger("номер единицы массы (целое число в диапазоне 1–5)")
	}

	y = util.NoNNumber("масса тела в этих единицах")

	switch {
	case x == 1:
		fmt.Printf("%v\n", "килограмм")
		fmt.Printf("%v\n", y)
	case x == 2:
		fmt.Printf("%v\n", "миллиграмм в килограмм")
		fmt.Printf("%v\n", y*0.00001)
	case x == 3:
		fmt.Printf("%v\n", "грамм в килограмм")
		fmt.Printf("%v\n", y*0.001)
	case x == 4:
		fmt.Printf("%v\n", "тонна в килограмм")
		fmt.Printf("%v\n", 1000*y)
	case x == 5:
		fmt.Printf("%v\n", "центнер в килограмм")
		fmt.Printf("%v\n", 100*y)
	}
}
