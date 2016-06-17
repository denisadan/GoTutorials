package main

import "fmt"

func pointers() {
	a := 10

	var pointerA *int
	pointerA = &a

	fmt.Printf("A este de tipul %T\n", a)
	fmt.Printf("A are adresa %x\n", &a)
	fmt.Printf("P pointeaza catre adresa %x\n", pointerA)
	fmt.Printf("P pointeaza catre adresa variabilei cu valoarea %d\n", *pointerA)
}

func pointerToPointer() {
	var a int
	var pointerToA *int
	var pointerToPointer **int

	a = 1000
	pointerToA = &a
	pointerToPointer = &pointerToA

	fmt.Printf("Valoarea lui a = %d\n", a)
	fmt.Printf("Valoarea de la pointerul catre a = %d\n", *pointerToA)
	fmt.Printf("aloarea pointerului pointerului catre a = %d\n", **pointerToPointer)

}

func simpleFunction() {
	const MAX int = 5
	a := [] int{1, 2, 3, 4, 5}
	for i := 0; i < MAX; i++ {
		fmt.Printf("\nvaloarea lui a[%d] = %d\n", i, a[i])
	}
}

func arrayOfPointers() {
	const MAX int = 5
	a := [] int{1, 2, 3, 4, 5}
	var pointer [MAX]*int

	for i := 0; i < MAX; i++ {
		pointer[i] = &a[i]
	}
	for i := 0; i < MAX; i++ {
		fmt.Printf("\nAdresa lui a[%d] = %d si valoarea lui a[%d] = %d\n", i, pointer[i], i, *pointer[i])
	}
}

func swap(a *int, b *int) {
	var aux int = *a
	aux = *a;
	*a = *b;
	*b = aux;
}
func main() {
	pointers()
	pointerToPointer()
	simpleFunction()
	arrayOfPointers()

	a := 5; b := 8
	fmt.Println("before: ", a, b)
	swap(&a, &b)
	fmt.Println("after:", a, b)

}
