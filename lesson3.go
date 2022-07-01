package main

import "fmt"

func main() {

	var age int8 = 10

	if(age >= 18){
		fmt.Println("can vote")
	} else {
		fmt.Println("nah")
	}

	//switch statement
	var color string = "Red"

	switch(color){
	case "Blue":
		fmt.Println("this is blue")
	case "Red":
		fmt.Println("this red")
	default:
		fmt.Println("choose color")
	
	
	}



	//For loops:

	for a := 0; a <= 10; a++ {
		if(a % 2 == 0){
			fmt.Println(a, "is an even number")
		}
	}

	// a different way of writing a for loops:
	var a int = 0
	for; a <= 10; {
		if(a % 3 == 0){
			fmt.Println(a, "is an odd number")
		}
		a++
	}




}

