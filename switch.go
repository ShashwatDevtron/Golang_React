package main

import (
	"fmt"
)
 
func main(){
	
	switch day := 3; day {
	case 1: fmt.Println("Monday")
	case 2: fmt.Println("Tuesday")
	case 3: fmt.Println("Wednesday")
	case 4: fmt.Println("Thursday")
	default: fmt.Println("Invalid")
	}
}