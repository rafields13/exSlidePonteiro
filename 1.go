// Escreva uma função swap que receba dois ponteiros para int como argumentos e troque os valores apontados por eles.

package main

import "fmt"

func swap(ptrUm *int, ptrDois *int) {
	*ptrUm, *ptrDois = *ptrDois, *ptrUm
}

func main() {
	x := 9
	var ptrUm = &x

	y := 7
	var ptrDois = &y

	fmt.Println("Valor apontado pelo primeiro ponteiro:", *ptrUm)
	fmt.Println("Valor apontado pelo segundo ponteiro:", *ptrDois)

	swap(ptrUm, ptrDois)
	fmt.Println("Valor apontado pelo primeiro ponteiro após a troca:", *ptrUm)
	fmt.Print("Valor apontado pelo segundo ponteiro após a troca: ", *ptrDois)
}
