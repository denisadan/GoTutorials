package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func multipleValues() (int, int) {
	return 4, 4
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, i := range nums {
		total += i
	}
	fmt.Println("\nsum:", total)

}
func main() {
	result := plus(5, 7)
	fmt.Println("5+7 =", result)

	result2 := plusplus(4, 5, 7)
	fmt.Println("4+5+7 =", result2)

	multiplerez1, multiplerez2 := multipleValues()
	fmt.Println(multiplerez1)
	fmt.Println(multiplerez2)

	_, d := multipleValues()
	fmt.Println(d)

	sum(1, 2)
	sum(1, 2, 3)
	myNumbers := []int{1, 2, 3, 4}
	sum(myNumbers...)

	var denisa string
	var denisa2 int
	denisa = "bla bla"
	fmt.Println(denisa)
	fmt.Printf("has type %T ", denisa)
	fmt.Println(denisa2)
	fmt.Printf("has type %T", denisa2)

	denisa3 := 5
	fmt.Println(denisa3)
	fmt.Printf("has type %T", denisa3)

}