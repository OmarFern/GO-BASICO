package main

import "fmt"

func main() {
	/*-----------array ----------> food=[0,1,2,3]----*/
	var food [4] string
	   	food[0] = "🍔"
	   	food[1] = "🍕"
	   	food[2] = "🌭"
	   	food[3] = "🍟" 
		fmt.Println(food)// run --->[🍔 🍕 🌭 🍟]
 /*----------------------------------------------------*/
	food_2:= [3] string{
		"🍔",
		"🌭",
		"🍕",} 
		fmt.Println(food_2)//run -->[🍔 🌭 🍕]
/*-----------------------------------------------------*/		

	var nums[3] int
	    nums[1] = 7
	    fmt.Println(nums)// [0 7 0]
}
