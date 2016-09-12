// Робот может перемещаться в четырех направлениях («С» — север,
// «З» — запад, «Ю» — юг, «В» — восток) и принимать три цифровые ко-
// манды: 0 — продолжать движение, 1 — поворот налево, –1 — поворот на-
// право. Дан символ C — исходное направление робота и целое число N —
// посланная ему команда. Вывести направление робота после выполнения
// полученной команды
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		s string
		y int
	)
	s = ioutil.Symbol("исходное направление робота")

	fmt.Println(s)

	y = -2
	for y < -1 || y > 1 {
		y = ioutil.Integer("команда")
	}

	switch {
	case s == "С": // север
		switch {
		case y == -1:
			fmt.Println("запад")
		case y == 0:
			fmt.Println("север")
		case y == 1:
			fmt.Println("восток")
		}
	case s == "З": // запад
		switch {
		case y == -1:
			fmt.Println("юг")
		case y == 0:
			fmt.Println("запад")
		case y == 1:
			fmt.Println("север")
		}
	case s == "Ю": // юг
		switch {
		case y == -1:
			fmt.Println("запад")
		case y == 0:
			fmt.Println("юг")
		case y == 1:
			fmt.Println("восток")
		}
	case s == "В": // восток
		switch {
		case y == -1:
			fmt.Println("север")
		case y == 0:
			fmt.Println("восток")
		case y == 1:
			fmt.Println("юг")
		}
	}
}
