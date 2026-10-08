package main

import "fmt"

func main() {
	fmt.Println("cli_calculator")
	fmt.Println("enter natural number")

	var sayı1 int 
	fmt.Scanln(&sayı1)
	fmt.Println("enter second natural number")
	var sayı2 int 
	fmt.Scanln(&sayı2)
	fmt.Println("What operation do you want to perform? ( only -,+,*,/ )")
	var ıslem string 
	fmt.Scanln(&ıslem)
	if ıslem == "+" {
		fmt.Println("aftermath : ", sayı1 + sayı2)
	}
	if ıslem == "-" {
		fmt.Println("aftermath :", sayı1 - sayı2)
	}
	if ıslem == "*" {
		fmt.Println("aftermath :", sayı1 * sayı2)
	}	
	if ıslem == "/" {
		if sayı2 == 0 {
			fmt.Println("aftermath : is not posible ... ")
		} else {
			fmt.Println("aftermath :", sayı1 / sayı2)
		}
	}
    if ıslem != "+" && ıslem != "-" && ıslem != "*" && ıslem != "/" {
    fmt.Println("Invalid operation .... some thing wrong with you? ")
}
}