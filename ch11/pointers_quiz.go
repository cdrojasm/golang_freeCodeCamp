package main

import "fmt"

func main() {
	var x int = 50
	var y *int = &x
	fmt.Println(x)
	fmt.Println(y)
	*y = 100 // we only are changed the value of x, so y still stores the memory adress of x variable
	fmt.Println(x)
	fmt.Println(y)
}
