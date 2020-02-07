package main

import "fmt"

func main(){

	// there are 3 types of switch cases go language provided..

	i:=5
	
	//normal switch case
	
	switch i {
		case 1:
		     fmt.Println("i is 1")
		case 3:
		     fmt.Println("i is 3")
		case 5:
		     fmt.Println("i is 5")
		default:
		     fmt.Println("i is not 1 or 3 or 5")
	}

	// cheack multiple values in one case..
	
	switch i {
		case 1,3,5:
		      fmt.Println("i is 1 or 3 or 5")
		default:
		      fmt.Println("i is other number ")
	}

	// check conditions using switch..
	
	switch {
		case i==1:
			fmt.Println("i is equels to 5")
		case i>3 :
			fmt.Println("i is grather than 3")
		case i<5:
			fmt.Println("i is less than 5")
		default:
			fmt.Println("i never match")
	}
}