// Дана масса M в килограммах. Используя операцию деления нацело,
// найти количество полных тонн в ней (1 тонна = 1000 кг)
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x int
	x = ioutil.Integer("Введите массу M в килограммах")
	fmt.Printf("количество полных тонн: %v (т)\n", x/1000)
}
