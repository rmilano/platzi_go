package main
import "fmt"

func main() {
	var numero int = 100
	numero2 := 200
	fmt.Println(numero, numero2)

	var numeroEntero = 10
	var numeroDoble = 20.5
	resultado := numeroEntero + int(numeroDoble)
	fmt.Println("Resultado de la suma:", resultado)

	var nombre string = "Rodolfo"
	apellido := "Milano"

	nombreCompleto := nombre + " " + apellido
	fmt.Println("Nombre completo:", nombreCompleto)

}

