// Даны два целых числа A и B ( A < B ). Найти сумму квадратов всех целых
// чисел от A до B включительно
package main

import (
	"fmt"

	"strconv"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var m, n int

	m = ioutil.Integer("A")
	for n < m {
		n = ioutil.Integer("B")
	}

	fmt.Println("сумму квадратов всех чисел")
	fmt.Println(foo1(m, n))
}

func foo1(m, n int) string {
	i := m
	m = m * m
	for i < n {
		i++
		m = m + i*i
	}
	s := strconv.Itoa(m)
	return s
}
