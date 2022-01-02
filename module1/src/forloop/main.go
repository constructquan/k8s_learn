package main

import (
	"fmt"
)

func main(){
	myArray := [5]string {"I", "am", "stupid", "and", "weak"}
	fmt.Printf("myArray %v\n", myArray )
      	for index,value  := range myArray{
		if value == "stupid" {
			myArray[index] = "smart"
		}
		if value == "weak" {
			myArray[index] = "strong"
		}
	}

	fmt.Printf("myArray %+v\n", myArray)
}
