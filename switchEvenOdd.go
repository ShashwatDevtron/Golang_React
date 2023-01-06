package main

import (
	"fmt"
)

func evenOdd(n int) string{
	
	flag  := false
	if(n %2 == 0){
		flag = true;
	}

	switch flag{
	case true: return "Number is even"
	case false: return "Number is odd"
	default : return "Number is Invalid"
	}
}

func main(){
	a := 9
	b := 10

	fmt.Println(evenOdd(a))
	fmt.Println(evenOdd(b))
	
}