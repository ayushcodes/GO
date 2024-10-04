package main

import "fmt"

func main(){

	a := make([]int, 5, 10)

	a[0] = 1
	a[1] = 2
	a[2] = 3

	a[3] = 4

	b := a[1:3]
	fmt.Println(a,b)
	b[0] = 342143
	fmt.Println(a,b)
	a[0] = 9000
	a[2] = 8000

	fmt.Println(a,b)
}