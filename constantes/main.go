package main

import "fmt"

func main() {
	var x int = 500
	var y int = -4500
	fmt.Printf("Type: %T, value: %v", x, x)
	fmt.Printf("\nType: %T, value: %v", y, y)

	variables()
  }
  /*-----------------------------------------------------------------*/
func variables(){
  var a bool = true     // Boolean
  var b int = 5         // Integer
  var c float32 = 3.14  // Floating point number
  var d string = "Hi!"  // String

  fmt.Println("\n Boolean: ", a)
  fmt.Println(" Integer: ", b)
  fmt.Println(" Float:   ", c)
  fmt.Println(" String:  ", d)



}