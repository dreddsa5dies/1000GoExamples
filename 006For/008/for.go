// Даны два целых числа A и B ( A < B ). Найти произведение всех целых чи-
// сел от A до B включительно
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

	fmt.Println("произведение")
	fmt.Println(foo1(m, n))
}

func foo1(m, n int) string {
	i := m + 1
	for i < n {
		m = m * i
		i++
	}
	s := strconv.Itoa(m * n)
	return s
}
