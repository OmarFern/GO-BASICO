package main


import "fmt"

func main() {
	/*------1---------*/
	condicional_1()
	/*------2--------*/
	condicional_2 ()
	/*------1---------*/
	edad() 
}
/*--------------------------------------------------------*/
func edad (){
		edad := 20
		if edad >= 18{
			fmt.Println("Mayor de edad")
		}
	}
/*--------------------------------------------------------*/
func condicional_1(){
	    emoji := "🐈"

	   	if emoji == "🌵" {
	   		fmt.Println("es un cactus")
	   	} else if emoji == "😋" {
	   		fmt.Println("es una carita")
	   	} else {
	   		fmt.Printf("el emoji es: %s\n", emoji)

}}
/*--------------------------------------------------------*/
func condicional_2 (){

	if emoji := "😋"; emoji == "🌵" {
		fmt.Println("es un cactus")
	} else if emoji == "😋" {
		fmt.Println("es una carita")
	} else {
		fmt.Printf("el emoji es: %s\n", emoji)
	} 
}