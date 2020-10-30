package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var arreglo []string
	var auxString string
	var temp int

	fmt.Print("Cuantas palabras vas a ingresar? ")
	fmt.Scan(&temp)

	for temp > 0 {
		fmt.Print("Ingresa Palabra: ")
		fmt.Scan(&auxString)
		arreglo = append(arreglo, auxString)
		temp --
	}

	fmt.Println("Arreglo desordenado:", arreglo)
	sort.Strings(arreglo)

	file, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	for i := 0; i < len(arreglo); i++ {
		file.WriteString(arreglo[i] + "\n")
	}

	file.Close()
	
	file, err = os.Create("descendente.txt")
	
	if err != nil {
		fmt.Println(err)
		return
	}

	for j := len(arreglo) - 1; j >= 0; j-- {
		file.WriteString(arreglo[j] + "\n")
	}

	file.Close()
}