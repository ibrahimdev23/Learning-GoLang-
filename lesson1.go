// In GoLang - a Go statement is a series
// of tokens. A token refers to any meaningful
// object in a statement like:
// identifiers, keywords, operators and punctuation
// and literals


//hello world
package main // package is a keyword
import(
	"fmt"
	"strconv"
)  //imports a libary 


// func main() {
// 	fmt.Println("Hello, World")
// }

//practice writing statements
func main() {
	fmt.Println("My name is Barbara")
	fmt.Println("I live in Augusta")
	fmt.Println("I like playing bball")
}

//In Go, a statement must end in a semicolon or in one of the following tokens:
// an identifier
// an integer, floating-point, imaginary, rune, or string literal
// one of the keywords break, continue, fallthrough, or return
// an operator or and punctuation, such as ++, --, ), ], or }

//hen you declare a variable, the var statement must include the data type for that variable.
// var x int = 20

//identifiers are tokens that identify things like functions, types, and variables. 
//identifiers can start with a letter or underscore 
//age_15 or Age or _age20

// Go is case sensitive 
//Go's name convention is either camleCase or PascalCase

func main() {
	fmt.Println("Enter your name")
	var firstName string
	
	fmt.Scanln(&firstName) //Scanln asks user for input and saves it to firstName
	fmt.Println(firstName)

	var address string
	// var zipcode int
	fmt.Println("enter address")
	fmt.Scanln(&address)
	// fmt.Scanln(&zipcode)

	fmt.Print("addrss is:" + " " + address)

	var lastName string = "Cena"  // this is a static type declaration 
	fmt.Println(lastName)
	var favNum = 29
	fmt.Print(&favNum)  // the & will access the memory location of the data
	
	//declaring multipe varibales:
	var message, email, location string
	message = "multiple variables"
	email = "jdkfj@jdi.com"
	location = "123 abc"
	fmt.Println(message, email, location)



// dynamic type declaration is where the complier will infer the type based on an assigned value
identifier := initial_value
	email := "john@cena.com"
	fmt.Print(email)
//cannot use := and var togther
var email := "djofj@goej.ocm" //this will throw an error
dyamic type also allows:
var message, year = "hello", 20

}


func main() {

	var firstNumber string 
	var secoundNumber string

	var firstInt int
	var secondInt int

	fmt.Print("Enter number")
	fmt.Scan(&firstNumber)
	fmt.Print("another number")
	fmt.Scan(&secoundNumber)

	firstInt, _ = strconv.Atoi(firstNumber) //strconv is a libary which converts str to int
	// secondInt, _ = strconv.Atoi(secoundNumber)
	// originalValue, error = strconv.Atio(integerValue)

	fmt.Println(firstInt + secondInt) //this adds the numbers total 


	// local variables only accessed within the block or function where defined

		// func main(){
		// this is a local variable 
		// 	var message string = "hello world"

		// }



}



