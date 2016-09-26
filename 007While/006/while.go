// Дано целое число N (> 0). Найти двойной факториал N :
// N !! = N ·( N –2)·( N –4)·...
// (последний сомножитель равен 2, если N — четное, и 1, если N — нечет-
// ное). Чтобы избежать целочисленного переполнения, вычислять это произ-
// ведение с помощью вещественной переменной и вывести его как вещест-
// венное число

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		n int
		x float64
	)

	for n <= 0 {
		n = ioutil.Integer("N")
	}
	x = countLoop(n)
	fmt.Printf("N !! = N ·( N –2)·( N –4)·... = %v\n", x)

}

func countLoop(n int) float64 {
	x := float64(n)
	y := 2.0
	if n%2 == 0 {
		for (float64(n) - y) != 2 {
			x = x * (x - y)
			y = y + 2
		}
	} else {
		for (float64(n) - y) != 1 {
			x = x * (x - y)
			y = y + 2
		}
	}
	return x
}
