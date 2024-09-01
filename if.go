package main

import "fmt"

func main(){
	a:=5
	if a< 0{
		fmt.Println("Negative Number")
	} else if a> 0{
		fmt.Println("Positive Number")
	} else{
		fmt.Println("Zero")
	}
}