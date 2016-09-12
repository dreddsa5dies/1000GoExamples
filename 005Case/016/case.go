// Дано целое число в диапазоне 20–69, определяющее возраст (в годах).
// Вывести строку-описание указанного возраста, обеспечив правильное со-
// гласование числа со словом «год», например: 20 — «двадцать лет», 32 —
// «тридцать два года», 41 — «сорок один год»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var m int

	for m < 20 || m > 69 {
		m = ioutil.Integer("число")
	}

	getYears(m)
}

func getYears(m int) {
	var w1, w2, w3 string
	switch {
	case (m / 10) == 2:
		w1 = "двадцать"
	case (m / 10) == 3:
		w1 = "тридцать"
	case (m / 10) == 4:
		w1 = "сорок"
	case (m / 10) == 5:
		w1 = "пятьдесят"
	case (m / 10) == 6:
		w1 = "шестьдесят"
	}

	switch {
	case m%10 == 1:
		w2 = "один"
	case m%10 == 2:
		w2 = "два"
	case m%10 == 3:
		w2 = "три"
	case m%10 == 4:
		w2 = "четыре"
	case m%10 == 5:
		w2 = "пять"
	case m%10 == 6:
		w2 = "шесть"
	case m%10 == 7:
		w2 = "семь"
	case m%10 == 8:
		w2 = "восемь"
	case m%10 == 9:
		w2 = "девять"
	}

	switch {
	case m%10 == 1:
		w3 = "год"
	case m%10 > 1 && m%10 < 5:
		w3 = "года"
	case m%10 > 4 && m%10 <= 9:
		w3 = "лет"
	}

	fmt.Printf("%v %v %v\n", w1, w2, w3)
}
