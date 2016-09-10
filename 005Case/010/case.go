// Робот может перемещаться в четырех направлениях («С» — север,
// «З» — запад, «Ю» — юг, «В» — восток) и принимать три цифровые ко-
// манды: 0 — продолжать движение, 1 — поворот налево, –1 — поворот на-
// право. Дан символ C — исходное направление робота и целое число N —
// посланная ему команда. Вывести направление робота после выполнения
// полученной команды
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var (
		x, y int
	)

	for x < 1 || x > 4 {
		x = util.Integer("исходное направление робота")
	}

	y = -2
	for y < -1 || y > 1 {
		y = util.Integer("команда")
	}

	switch {
	case x == 1: // север
		switch {
		case y == -1:
			fmt.Println("запад")
		case y == 0:
			fmt.Println("север")
		case y == 1:
			fmt.Println("восток")
		}
	case x == 2: // запад
		switch {
		case y == -1:
			fmt.Println("юг")
		case y == 0:
			fmt.Println("запад")
		case y == 1:
			fmt.Println("север")
		}
	case x == 3: // юг
		switch {
		case y == -1:
			fmt.Println("запад")
		case y == 0:
			fmt.Println("юг")
		case y == 1:
			fmt.Println("восток")
		}
	case x == 4: // восток
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
