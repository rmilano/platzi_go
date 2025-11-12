package main	
import "fmt"
 
func main() {

	i := 1
	for i <= 3 { 
		fmt.Println(i)
		i++ // puedes hacer esto o puedes usar i += 1,  i = i + 1
		/*
		aqui puede poner mucho texto
		*/
	}

	fmt.Println("--------------------")

	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}

	fmt.Println("--------------------")

	for {
		fmt.Println("loop infinito")
		break
	}

	fmt.Println("--------------------")

	for valor := range 6 {
		fmt.Println(valor)
		if valor%2 == 0 {
			continue
		}	
		fmt.Println(valor)
	}
}
