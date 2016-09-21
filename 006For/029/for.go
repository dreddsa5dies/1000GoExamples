// Дано целое число N (> 1) и две вещественные точки на числовой оси:
// A , B ( A < B ). Отрезок [ A , B ] разбит на N равных отрезков. Вывести H —
// длину каждого отрезка, а также набор точек
// A , A + H , A + 2· H , A + 3· H , ..., B
// образующий разбиение отрезка [ A , B ]

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		a, b float64
		n    int
	)

	n = -1
	for n <= 1 {
		n = ioutil.Integer("N")
	}

	a = ioutil.NotNullNumber("A")
	b = a - 1
	for a > b {
		b = ioutil.NotNullNumber("B")
	}

	fmt.Println("A , A + H , A + 2· H , A + 3· H , ..., B ,")
	x := foo1(n, a, b)
	i := a
	for i <= b {
		fmt.Printf("%v\t", i)
		i = i + x
	}
	fmt.Println()
}

func foo1(n int, a, b float64) float64 {
	k := (b - a) / float64(n)
	return k
}
