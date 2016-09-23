// Даны целые положительные числа A и B ( A < B ). Вывести все целые
// числа от A до B включительно; при этом каждое число должно выводиться
// столько раз, каково его значение (например, число 3 выводится 3 раза)

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		a, b uint
	)

	for a <= 0 {
		a = ioutil.UInteger("A")
	}

	b = a - 1
	for b < a {
		b = ioutil.UInteger("B")
	}
	foo1(a, b)
}

func foo1(a, b uint) {
	var i, l uint
	for i = a; i <= b; i++ {
		for l = 1; l <= i; l++ {
			foo2(i)
		}
		fmt.Println()
	}
}

func foo2(i uint) {
	fmt.Printf("%v\t", i)
}
