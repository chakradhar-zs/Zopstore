package main

import "fmt"

// Write a Golang program to implement a function factorial using
// function closures to calculate the factorial of a given positive integer.

func factorial() func(int) int {

}

func main() {
	f := factorial()
	fmt.Println(f(5)) // prints 120
}
