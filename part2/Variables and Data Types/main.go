package main

import "fmt"

func main() {
	//In Go, to declare a variable, we first use the `var` keyword, then specify the variable's name, and then its type. Next, we place the `=` sign and write the variable's value inside the "" quotes.
	
	// for letter use string
	var name string = "hackpatato"
	// If we want to print variables, we don't need to use quotation marks after opening the parentheses in fmt.Println
	fmt.Println(name)
	// NOTE: In the Go programming language, we don't always need to specify a variable's type. Go can figure this out on its own.
	var surname = "secret"
	fmt.Println(surname)


	//When working with numbers, we don't need to use the “” symbol. We use `int` for integers and `float` for numbers with decimals.
	var number1 int = 13
	var float1 float32 = 2.4892
	var float2 = 288.24812941
	fmt.Println(number1)
	fmt.Println(float1)
	println(float2)

	var isActive bool = true
	var İsfalse bool = false
	fmt.Println(isActive)
	fmt.Println(İsfalse)
	// A boolean is a variable that indicates whether a given variable is false or true. It is either true or false.
	var a, b, c int = 1, 2, 3
	// Yes, we can declare variables in multiple ways on a single line.
	fmt.Println(a, b, c)
}