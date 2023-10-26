package main

import "fmt"

func main() {
	//Поменять местами два числа без создания временной переменной.
	//Если имеется в виду поменять значения в переменных, то:
	x, y := 5, 10
	x = x ^ y
	y = y ^ x
	x = x ^ y
	fmt.Println(x, y)

	x = x + y
	y = x - y
	x = x - y
	fmt.Println(x, y)
}
