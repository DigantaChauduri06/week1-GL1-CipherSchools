package main

import "fmt"


func Xyz() {
	// fmt.Println("Enter a number")
	// var input int
	// fmt.Scanln(&input)
	// if input & 1 != 0 {
	// 	fmt.Println("It is odd")
	// 	} else {
	// 	fmt.Println("It is even")

	// }

	// if  x:= 10; x == 10 {
	// 	fmt.Println("x is 10")
	// }
	switch code := 100; code {
		case 100:
			fmt.Println("100")
			fallthrough
		case 200:
			fmt.Println("100")
			fallthrough
		case 300:
			fmt.Println("100")
			fallthrough
		case 400:
			fmt.Println("100")
			fallthrough
		case 500:
			fmt.Println("100")
			fallthrough
		case 600:
			fmt.Println("100")
			fallthrough
		case 700:
			fmt.Println("Fallen through")
	}

}