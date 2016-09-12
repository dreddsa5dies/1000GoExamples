// Даны два целых числа: D (день) и M (месяц), определяющие правиль-
// ную дату. Вывести знак Зодиака, соответствующий этой дате: «Водолей»
// (20.1–18.2), «Рыбы» (19.2–20.3), «Овен» (21.3–19.4), «Телец» (20.4–20.5),
// «Близнецы» (21.5–21.6), «Рак» (22.6–22.7), «Лев» (23.7–22.8), «Дева»
// (23.8–22.9), «Весы» (23.9–22.10), «Скорпион» (23.10–22.11), «Стрелец»
// (23.11–21.12), «Козерог» (22.12–19.1)
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var m, n int
	for m < 1 || m > 12 {
		m = ioutil.Integer("месяц")
	}

	if m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10 || m == 12 {
		for n < 1 || n > 31 {
			n = ioutil.Integer("день")
		}
	} else if m == 4 || m == 6 || m == 9 || m == 11 {
		for n < 1 || n > 30 {
			n = ioutil.Integer("день")
		}
	} else if m == 2 {
		for n < 1 || n > 28 {
			n = ioutil.Integer("день")
		}
	}

	fmt.Printf("%s\n", printZodiac(m, n))
}

func printZodiac(m, n int) string {
	var x int
	w := [...]string{1: "Водолей",
		2:  "Рыбы",
		3:  "Овен",
		4:  "Телец",
		5:  "Близнецы",
		6:  "Рак",
		7:  "Лев",
		8:  "Дева",
		9:  "Весы",
		10: "Скорпион",
		11: "Стрелец",
		12: "Козерог"}

	switch {
	case m == 1:
		switch {
		case n >= 20:
			x = 1
		case n < 20:
			x = 12
		}
	case m == 2:
		switch {
		case n <= 18:
			x = m - 1
		case n > 18:
			x = m
		}
	case m == 3:
		switch {
		case n <= 20:
			x = m - 1
		case n > 20:
			x = m
		}
	case m == 4:
		switch {
		case n <= 19:
			x = m - 1
		case n > 19:
			x = m
		}
	case m == 5:
		switch {
		case n <= 20:
			x = m - 1
		case n > 20:
			x = m
		}
	case m == 6:
		switch {
		case n <= 21:
			x = m - 1
		case n > 21:
			x = m
		}
	case m == 7:
		switch {
		case n <= 22:
			x = m - 1
		case n > 22:
			x = m
		}
	case m == 8:
		switch {
		case n <= 22:
			x = m - 1
		case n > 22:
			x = m
		}
	case m == 9:
		switch {
		case n <= 22:
			x = m - 1
		case n > 22:
			x = m
		}
	case m == 10:
		switch {
		case n <= 22:
			x = m - 1
		case n > 22:
			x = m
		}
	case m == 11:
		switch {
		case n <= 22:
			x = m - 1
		case n > 22:
			x = m
		}
	case m == 12:
		switch {
		case n <= 22:
			x = m - 1
		case n > 22:
			x = m
		}
	}
	return w[x]
}
