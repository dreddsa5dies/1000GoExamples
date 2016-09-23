// Дано целое число N (> 1) и две вещественные точки на числовой оси:
// A , B ( A < B ). Отрезок [ A , B ] разбит на N равных отрезков. Вывести H —
// длину каждого отрезка, а также значения функции F ( X ) = 1 – sin( X ) в точ-
// ках, разбивающих отрезок [ A , B ]:
// F ( A ), F ( A + H ), F ( A + 2· H ), ..., F ( B )

package main

import (
	"fmt"
	"math"

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

	fmt.Println("F ( A ), F ( A + H ), F ( A + 2· H ), ..., F ( B )")
	fmt.Printf("Длина отрезка = %v\n", foo1(n, a, b))
	foo2(n, a, b)
}

// расчет длины отрезка
func foo1(n int, a, b float64) float64 {
	k := (b - a) / float64(n)
	return k
}

// печать 1-sin(X) отрезка
func foo2(n int, a, b float64) {
	x := foo1(n, a, b)
	i := a
	for i <= b {
		if (b - i) > 0.0001 {
			fmt.Printf("1-sin(%v) = %v\n", i, 1-math.Sin(i))
		} else {
			fmt.Printf("1-sin(%v) = %v\n", b, 1-math.Sin(b))
		}
		i = i + x
	}
}
