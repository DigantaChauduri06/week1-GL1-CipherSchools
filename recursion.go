package main

import "fmt"

func Rec(num int) {
	if (num <= 0) {
		return
	}
	if (num %2 == 0) {
		fmt.Println(num + 1)
		Rec(num-2)
		} else {
		fmt.Println(num + 2)
		Rec(num-1)
			
	}
	fmt.Println(num - 2)
}

func Fact(number int) int {
	if number == 1 || number == 0 {
		return 1
	}
	if number < 0 {
		fmt.Println("Invalid Number")
		return -1
	}
	res := number * Fact(number - 1)
	return res
}

func fib(n int64) int64 {
	if n <= 1 {
		return n
	}
	return fib(n - 1) + fib(n - 2)
}

func _main() {
	fmt.Println(fib(4))
}