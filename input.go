package main

import(
	"fmt"
	"math"
)

func compoundIntrest(p, r, t, n int)float64{
	 var a float64 =  math.Pow((float64)(p*(1+(r/n))),(float64)(n*t))
	return a
}

func main(){
	var p, r,t, n int
	fmt.Println("Enter the principal")
	fmt.Scan(&p)
	fmt.Println("Enter the rate")
	fmt.Scan(&r)
	fmt.Println("Enter the time")
	fmt.Scan(&t)
	fmt.Println("Enter the number of times the intrest will be compounded")
	fmt.Scan(&n)
	fmt.Println(compoundIntrest(p,r,t,n))
}