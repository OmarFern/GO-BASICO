package main

import "fmt"


func main() {
	fmt.Println("-------------bucle for 0-----------------")
	bucle_0()
	fmt.Println("-------------bucle for 1-----------------")
	bucle_1()
	fmt.Println("-------------bucle for 2-----------------")
	bucle_2()
	fmt.Println("-------------bucle for 3-----------------")
	bucle_3()
	fmt.Println("-------------bucle for 4-----------------")
	bucle_4()
	fmt.Println("-------------bucle for 5-----------------")
	bucle_5()
	fmt.Println("-------------bucle for 6-----------------")
	bucle_6()
	fmt.Println("-------------bucle for 7-----------------")
	bucle_7()
	fmt.Println("-------------bucle for 8-----------------")
	bucle_8()
	fmt.Println("-------------bucle for 9-----------------")
	bucle_9() 
	fmt.Println("-------------bucle for 10----------------")
	bucle_10()
	fmt.Println("-------------bucle for 11----------------")
	bucle_11()
}
/*------------------------------*/
func bucle_0(){
	hello := "hello"
	for _, v := range hello {
		fmt.Println(string(v))
	}
}
/*------------------------------*/
func bucle_1(){
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
/*------------------------------*/
func bucle_2(){
	for i := 0; i < 15; i+=3 {
		fmt.Println(i)
	}
}
/*-------------------------------*/
func bucle_3(){
	for i := 50; i > 0; i -= 10 {
		fmt.Println(i)
}}
/*------------------------------*/
func bucle_4(){
	i := 0          //inicia
    for i < 5 {     //codicion
	fmt.Println(i)
	i++            //construccion
}
}
/*--------------------------------*/
func bucle_5(){
	
		sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}
	
		for i := 0; i < len(sharks); i++ {
			fmt.Println(sharks[i])
		}
}
/*-------------------------------*/
func bucle_6(){
	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

	for i, shark := range sharks {
		fmt.Println(i, shark)
	}
}
/*-------------------------------*/
func bucle_7(){
	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

	for range sharks {
		sharks = append(sharks, "shark")
	}

	fmt.Printf("%q\n", sharks)

}
/*--------------------------------*/
func bucle_8(){
	sammy := "Sammy"

	for _, letter := range sammy {
		fmt.Printf("%c\n", letter)
	}

}
/*---------------------------------*/
func bucle_9() {
	sammyShark := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}

	for key, value := range sammyShark {
		fmt.Println(key + ": " + value)
	}
}
/*-----------------------------------*/
func bucle_10() {
	ints := [][]int{
		{0, 1, 2},
		{-1, -2, -3},
		{9, 8, 7},
	}

	for _, i := range ints {
		fmt.Println(i)
	}
}
/*----------------------------------*/
func bucle_11() {
	ints := [][]int{
		{0, 1, 2},
		{-1, -2, -3},
		{9, 8, 7},
	}

	for _, i := range ints {
		for _, j := range i {
			fmt.Println(j)
		}
	}
}