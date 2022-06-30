package main

import (
	"fmt"

)


//TYPES
//once a type is assigned to a variable it cannot inherently change
// in go, type is assigned to the variable and not the value 



func main() {

	// quantity := 100 //go assings this variable as an int type
	// fmt.Printf("%T\n", quantity) // this prints the type 


	// Go does not allow calculations using different data types
	// if when the types are similar like floats and ints 
	//

	firstNum := 1.0
	var sNumm int = 20

	//  var total float32
	 var multiply float32
	//total = firstNum + sNumm  //this doesn't work b/c diffrent types
	//fmt.Print(add)

	// type casting is used to overcome this problem
	//total = int(firstNum) + int(sNumm) //state the type you want to convert to
	multiply = float32(sNumm) * float32(firstNum)
	fmt.Println(multiply)

	
	fmt.Printf("%T\n", sNumm) //the orginal varibales type is the same 
	


	// Relational Operations: the result is either true or false 
	// == checks if values are equal or not, if yes returns true 
	// e != checks if the not equal, returns true 
	// > 
	//< 
	// >= and <=

	//cannot compare variables of different types 


	//Assigment Operators
	// = , +=, -= , *= , /= , \%=


	//MATH package 
	// import "math"
	// includes Abs, Ceil, Floor, Exp, Sqrt, Trunc, Round, Round(a,b), Pow(a,b)


	// Bitwise Operators
	// & , | , 
	// ^ binary XOR operator, checks if bit is set in one operand but not both
	// << binary left shift operator, checks 
	// >> 


	//RANDOM NUMBERS 
	// use the "math/rand" package
	//floats are between 0 and 1
	//random int any integer over 0 unless specified
	// rand.Float32()
	//rand.Float64
	//rand.Int()
	//use Intn to limit range 
	//rand.Int(100) //this limit highest possible value to 100

	//there is a way to change the seeding so the rand method generates a random number everyime

	//Boolean Operations
	// &&, ||, ! 
	
	



}