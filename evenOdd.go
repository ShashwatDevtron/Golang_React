package main

import(
	"fmt"
)

func check(n int) string{
	if (n %  2 ==0){
		return "Number is Even"
	}else{
		return "Number is Odd"
	}
}

func main(){
	a := 2
	b := 3

	fmt.Println(check(a), check(b))
	
}