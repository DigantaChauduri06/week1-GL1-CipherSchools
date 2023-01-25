package main

import "fmt"
func main2() {
	// i:= 10
	// j := &i
	// *j = 100
	// fmt.Println(*j)
	// var i *int
	// *i = 10
	// j := &i
	// **j = 100
	// j:= new(int)
	name:= new(string)
	*name ="golang"
	fmt.Println(name)
	
}