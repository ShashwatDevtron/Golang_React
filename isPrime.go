package main

import(
	"fmt"
)

func isPrimeNumber(n int) string{
	for i :=2 ; i*i< n; i++{
		if(n % i == 0){
			return "The number is not Prime"
		}
	}
	return "The Number is Prime"
}

func main(){
	fmt.Println(isPrimeNumber(7))
	fmt.Println(isPrimeNumber(10))	
}