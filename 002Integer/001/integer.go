// Дано расстояние L в сантиметрах. Используя операцию деления наце-
// ло, найти количество полных метров в нем (1 метр = 100 см)
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x int
	x = util.Integer("Введите расстояние L в сантиметрах")
	fmt.Printf("количество полных метров: %v (м)\n", x/100)
}
