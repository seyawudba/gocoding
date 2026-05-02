package main

import "fmt"

var funcVar func(int) int

func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

func incFn(x int) int {
	return x + 1
}
func decFn(x int) int {
	return x - 1
}

func main() {

	fmt.Println(applyIt(incFn, 2))
	fmt.Println(applyIt(decFn, 2))

	v := applyIt(func(x int) int { return x + 1 }, 2)
	fmt.Println(v)
}
