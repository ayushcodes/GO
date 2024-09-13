package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":"r",
		"green":"3268476234",
	}
	fmt.Printf("%+v", colors)
	delete(colors, "red")
	var test map[string]string
	test2 := make(map[string]string)
	fmt.Println(test,test2)
}