package main

import (
	"fmt"
)

func main() {
	var celsius float64
	fmt.Println("Informe a temperatura em graus Celsius!")
	fmt.Scanf("%f", &celsius)
	conversor(celsius)
}

func conversor(grau float64) {
	fahrenheit := grau*9/5 + 32
	fmt.Println(grau, "graus celsius é igual a", fahrenheit, "em Fahrenheit ")
}
