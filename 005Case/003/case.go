// Дан номер месяца — целое число в диапазоне 1–12 (1 — январь, 2 —
// февраль и т. д.). Вывести название соответствующего времени года («зи-
// ма», «весна», «лето», «осень»)
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x int
	for x < 1 || x > 12 {
		x = ioutil.Integer("число")
	}

	switch {
	case x == 3 || x == 4 || x == 5:
		fmt.Println("весна")
	case x == 6 || x == 7 || x == 8:
		fmt.Println("лето")
	case x == 9 || x == 10 || x == 11:
		fmt.Println("осень")
	case x == 1 || x == 2 || x == 12:
		fmt.Println("зима")
	}
}
