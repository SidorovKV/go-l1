package main

import "fmt"

func main() {
	//Поменять местами два числа без создания временной переменной.
	//Если имеется в виду поменять в переменных, то всё просто:
	x, y := 5, 10
	x, y = y, x
	fmt.Println(x, y)
	//Можно и функцией:
	x, y = swap(x, y)
	fmt.Println(x, y)
	x, y = swap(x, y)
	fmt.Println(x, y)

	//Если в массиве, вроде тоже не сложно:
	a := [4]int{1, 2, 5, 8}
	a[0], a[2] = a[2], a[0]
	fmt.Println(a)
	swapInArray(a[:], 0, 2)
	fmt.Println(a)
}

func swap(x, y int) (int, int) {
	return y, x
}

func swapInArray(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}
