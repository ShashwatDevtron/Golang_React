package main

import (
	"fmt"
	
)

func secondLargestElement(a,b,c,d int)int{
	var largest int = a 
	var secondLarget int = b

	if(b>largest){
		secondLarget = largest
		largest = b
		
	}
	if(c>largest){
		secondLarget = largest
		largest =c
	}
	if(d>largest){
		secondLarget = largest
		largest = d
	}
	return secondLarget
}

func main(){

	var a,b,c,d int
	fmt.Println("Enter 1st number")
	fmt.Scanln(&a)
	fmt.Println("Enter 2nd number")
	fmt.Scanln(&b)
	fmt.Println("Enter 3rd number")
	fmt.Scanln(&c)
	fmt.Println("Enter 4th number")
	fmt.Scanln(&d)

	fmt.Println(secondLargestElement(a,b,c,d))
}