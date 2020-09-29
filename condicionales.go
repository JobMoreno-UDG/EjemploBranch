package main

import "fmt"

func main(){
	//Condicionales
	var temp int
	fmt.Print("Temp: ")
	fmt.Scan(&temp)
	if temp>0{
		fmt.Println("Temperatura Positiva")
	}else if temp ==0 {
		fmt.Print("temperatura cero")
	}else {
		fmt.Println("Temperatura bajo cero")
	}
}