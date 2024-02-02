package main

import (
	"fmt"
	"projectTest/calc"
)

func main(){
	num1 := 4
	num2 := 2
	
	result := calc.Division(num1, num2)

	fmt.Println(result)
}