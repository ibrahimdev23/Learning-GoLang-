package main

import "fmt"

func main() {



//
// arr1 := [][]int{}

// 	row1 := []int{}
// 	row2 := []int{}
// 	row3 := []int{}

// 	var num2 int
// 	var num3 int
// 	var num4 int

// 	for i := 0; i < 3; i++{
// 		fmt.Scan(&num2)
// 		row1 = append(row1, num2)
// 	}

// 	for i := 0; i < 3; i++{
// 		fmt.Scan(&num3)
// 		row2 = append(row2, num3)
// 	}
// 	for i := 0; i < 3; i++{
// 		fmt.Scan(&num4)
// 		row3 = append(row3, num4)
// 	}
	
// 	arr1 = append(arr1, row1)
// 	arr1 = append(arr1, row2)
// 	arr1 = append(arr1, row3)

	
// 	fmt.Println(arr1)
	

// }





var arr1 [][] int

	for i := 0; i < 3; i++{
		arr1 = append(arr1, userInputRow())
	}
	fmt.Println(arr1)


}
//more optimized way using one function to get user input 
	func userInputRow() []int{
		rows := make([]int, 3)
	
		for i := 0; i < len(rows); i++{
			var num int 
			fmt.Scan(&num)
			rows[i] = num
		}
		return rows
	}
	
