package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func main() {
	result := sum(3, 5)
	fmt.Println("The sum is:", result)
}
