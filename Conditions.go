package main

import "fmt"

func main(){
	
	// if , and else conditions in go..
	// ternory operator not provides to go .such as a?b:c
	i:=5

	if i<5{
		fmt.Println(i," is less than 5")
	}else if i==5{
		fmt.Println(i," is equels to 5")
	}else{
		fmt.Println(i," is grather than 5")
	}

	for i:=0;i<=10;i++{
		
	    for j:=0;j<=i;j++{
		if j%2 ==0 {
		   fmt.Print("*")
		}else{
		   fmt.Print("o")
		}
	
	    }
	    fmt.Println()
	}

	
}