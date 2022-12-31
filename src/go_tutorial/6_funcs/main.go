package main

import (
	"fmt"
	"6_funcs/localpackage"
)

type User struct {
	name string
	age int
}

func main(){
	var name string = "Satoru"
	var age int = 28

	user := second.ConvertUser(name,age)

	if second.PrintlnUser(user) {
		fmt.Println("^^^")
	}
}