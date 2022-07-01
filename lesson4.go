package main
import "fmt"

//Functions
// basic function syntax:
// func funcName (arg1 type, arg2 type) return_type {
//	[function instructions ]
//}

func add(a int, b int) int {
	return a + b
}

// shorter way of writing the same:
func add2(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("results:", add(5, 5))

}
//Functions with no return value 



//Functions as a value: we can define a function as a variable and then reference it when needed
//
// add := func() int {
//	return a + b
//}

// 	supermutation := func(a,b,c int) (x,y,z int){
// 		x = a + b
// 		y = a + c
// 		z = b + c
// 		return 
// 	}
// 	fmt.Println((supermutation(4,6,5)))
// }



//Variadic function is a function that accepts a varible number of inpus
// func funcNmae (para....type) [return_type] {
//	function instructions
//}


//


//Recurion

 func factorial(n int) int {
	if n == 0{
		return 1
	}
	return n * factorial(n-1)
 }

func dis(){
	
	fmt.Println(factorial(5))
}

// //