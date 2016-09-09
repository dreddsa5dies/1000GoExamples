// Единицы длины пронумерованы следующим образом: 1 — дециметр,
// 2 — километр, 3 — метр, 4 — миллиметр, 5 — сантиметр. Дан номер еди-
// ницы длины (целое число в диапазоне 1–5) и длина отрезка в этих едини-
// цах (вещественное число). Найти длину отрезка в метрах
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
		x = util.UInteger("номер единицы длины (целое число в диапазоне 1–5)")
	}

	y = util.NoNNumber("длина отрезка")

	switch {
	case x == 1:
		fmt.Printf("%v\n", "дециметр в метр")
		fmt.Printf("%v\n", 0.1*y)
	case x == 2:
		fmt.Printf("%v\n", "километр в метр")
		fmt.Printf("%v\n", y*1000)
	case x == 3:
		fmt.Printf("%v\n", "метр")
		fmt.Printf("%v\n", y)
	case x == 4:
		fmt.Printf("%v\n", "миллиметр в метр")
		fmt.Printf("%v\n", 0.001*y)
	case x == 5:
		fmt.Printf("%v\n", "сантиметр в метр")
		fmt.Printf("%v\n", 0.01*y)
	}
}
