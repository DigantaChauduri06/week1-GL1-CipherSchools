package main

import "fmt"

func Main_() int {
	member := 10
	var num int 
	var function2 func(a int) int
	function := func () int {
		num = 100
		fmt.Println("In a func", num)
		return member * 100
	}
	
	functi := func (a int) int {
		return a
	}(100)
	function2(100)
	fmt.Println(functi)
	g(function)

	var y func()
	y = func ()  {
		
	}
	y()
	return function()
}

func g(getInt func() int) string {
	return "Yay"
}