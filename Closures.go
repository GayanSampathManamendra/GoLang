package main

import "fmt"

func hello_function(msg string)string{
	return msg
}

func retur_anonymous_function() func(string){
	
	return func(msg string){
		fmt.Println(msg)
	}
}

func int_seq() func()int{

	i := 0
	return func()int{
	   i++
	   return i
	}
}

func main(){

	fmt.Println("This is main function ...")

	fmt.Println(hello_function("hello world !"))

	// anonymous function declared..
	
	func(msg string){
		fmt.Println(msg)
	}("This is anonymous function !")

	print_anonymous := retur_anonymous_function()
	print_anonymous("This function return an another anonymous function !")

	seq1 := int_seq();

	// in this only return the memory addres where is the anonymous function stored..
	fmt.Println(seq1)
	
	// to get the return value of anonymous function ..
	fmt.Println(seq1())
	fmt.Println(seq1())

	// create a new varible and assigned the int_seq() function
	seq2 := int_seq()
	fmt.Println(seq2)
	fmt.Println(seq2())
}



