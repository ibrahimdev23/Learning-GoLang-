package main

import (
	"fmt"
)



//A hash map is a collection of paired values, where one value in each pair is a key that identifies
//  the other value in the pair. We can use maps to reference values through 
//  their key, as a way of avoiding referencing the value itself, or 
//  we can use maps to change one value that is associated with another value.


//Creating a map:
 //   map[KeyType]ValueType
 //map[string]int meaning {"string" : int}


// 	freqCounter := map[string]int{
// 		"hi": 23,
// 		"hello": 2,
// 		"hey": 4,
// 		"weather": 1,
// 		"greet": 35,
// 	 }
  
// 	 fmt.Println("Map value:", freqCounter)
//   



type class struct {
	ClassName string
	students []student //student is data type 
}


  type student struct{
	Name string 
	RollNo int
	City string
  }



func main(){


	abs := student{Name:"ross", RollNo: 1, City: "NewYork"}
	cbs := student{Name:"James", RollNo: 2, City: "Chicago"}


	// students := student{abs,cbs}
	students:= []student{abs,cbs, student{Name:"Jack", RollNo:3, City: "London"}}

	//slice of slice of struc 
	class := class{"firstA", students}
	fmt.Println(class)

	for i, j:=range students{
		fmt.Println(i, " ", j)
	}
  }



  