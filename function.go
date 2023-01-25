package main

import "fmt"

/*
	func hello() {
		----
		----
		----
	}

*/


func a() {
	fmt.Print("Calling a")
}
func _b(args ...int) int {
	a()
	return 100
}
func _c(a *int,b *string,c *string) string{
	a_val := _b(10,2,3,4)
	return fmt.Sprint(a_val)
}
func d() int{
	a:=10
	b:="Name"
	c:="Age"
	// a:=10
	// a:=10
	_c(&a, &b, &c)
	return 1092

}
func e() (int, string, bool){
	d()
	return 10,"sdfsd",true
}