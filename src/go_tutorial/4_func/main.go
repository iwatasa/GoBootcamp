package main

import "fmt"

func myPrintln(text string) {
	fmt.Println(text)
}

func main(){
	var text string = "Hello World!!"
	myPrintln(text)
}