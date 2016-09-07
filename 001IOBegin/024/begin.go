// Даны переменные A, B, C. Изменить их значения, переместив содер-
// жимое A в C, C — в B, B — в A, и вывести новые значения переменных A,
// B, C.
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y, z float64
	x = util.Number("Введите A")
	y = util.Number("Введите B")
	z = util.Number("Введите C")
	x = x + z
	z = x - z // это теперь А
	x = x - z // это теперь C
	x = x + y
	y = x - y // это теперь C
	x = x - y // это теперь B
	fmt.Println("замена А\t: ", x)
	fmt.Println("замена B\t: ", y)
	fmt.Println("замена C\t: ", z)
}
