package main

import (
	"bufio"
	"log"
	"fmt"
	"os"
)
/*-------------------------------------------------------------------*/
func main() {
	file, err := os.Create("hello.txt")// creo el archivo.txt
	if err != nil {
		fmt.Printf("Ocurrió un error al crear el archivo: %v", err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte("Hello omar estoy manipulando un txt"))// escribo en el archivo.txt



	if err != nil {
		fmt.Printf("Ocurrió un error al escribir el archivo: %v", err)
		return
	}
	leer()
}
/*---------------------------------------------------------------------*/
func leer(){

	file, err := os.Open( "hello.txt")// una vez escrito lo abro el archivo 
	if err != nil {
	  log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
	  fmt.Println(scanner.Text())
	  fmt.Println(scanner.Bytes())
	}

}