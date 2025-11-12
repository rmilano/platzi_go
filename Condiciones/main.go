package main

import "fmt"

func main() {
	nombre := "Rodolfo"
	edad := 53

	if nombre == "Rodolfo" {
		fmt.Println("Hola Rodolfo")
	} else {
		fmt.Println("Hola Desconocido")
	}

	if edad >= 18 {
		fmt.Println("Ya puedes votar!")
	} else {
		fmt.Println("Eres menor de edad")
	}

	if 8%2 == 0 {
		fmt.Println("8 es un número par")
	} else {
		fmt.Println("8 es un número impar")
	}	

	if numero := 9; numero < 0 {
		fmt.Println(numero, "es negativo")
	} else if numero < 10 {
		fmt.Println(numero, "es un digito")
	} else {
		fmt.Println(numero, "es un numero mayor a 9")
	}	
}