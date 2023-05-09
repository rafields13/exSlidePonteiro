// Escreva uma função que receba um ponteiro para um booleano e altere o valor desse booleano para o seu inverso.

package main

import "fmt"

func reverseBool(ptr *bool) {
	*ptr = !*ptr
}

func main() {
	x := true
	var ptr = &x

	fmt.Println("Valor apontado pelo ponteiro:", *ptr)

	reverseBool(ptr)
	fmt.Print("Valor apontado pelo ponteiro após a inversão: ", *ptr)
}
