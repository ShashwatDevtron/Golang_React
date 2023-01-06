package main

import (
	"fmt"
)

func calc(n int) int{
	var count int
	for (n != 0){
		n = n/10;
		count++
	}
	return count
}

func main(){
	var n int
	fmt.Println("Enter a number")
	fmt.Scan(&n)
	result := calc(n)
	fmt.Println("The number of digits in number is", result)
}