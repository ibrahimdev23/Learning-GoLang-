package main
import "fmt"



//in Go, array values must be the same data type
//the size is part of the array's type
//once defined, array size cannot change but individual elements can be altered 
// if array is made up of int then the initial/default values are 0
// if its a string array then initial/defualt values are null

//creates an array with 5 elements of int type
var numbers [3] int

//assign a vlue to each positon in the array
// numbers[0] = 1
// numbers[1] = 2
// numbers[2] = 3


//more common way of creating an array
//declare an array and initialize its values
var nums = [5] int {5,4,3,2,1}

//or useing :=
// bigNums := [3] int {1,2,3}

//without var keyword
// words := [1] string {"hi, "yo"}


//inference of array size
//leaving out size then Go will create array equal to elements 
// numbers := [...] int {1,2,3,4,5,6,7}


//use for loop or range to ilerate over the array 

//Slices 
// a[low: high] //low is included but high is not
// primes := [6]int{2,3,5,7,11,13}
// var s []int = primes[1:4]   //this will made an array of [3,5,7]


//