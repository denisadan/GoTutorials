package main

import "fmt"

func simpleArrays() {
	var vector [5] int
	var matrix [5][5] int

	vector2 := [5] int{1, 2, 3, 4, 5}

	var denisachar string = "denisa"
	fmt.Println("lunigme:", len(denisachar))

	fmt.Println("vector:", vector)
	fmt.Println("matrix:", matrix)
	fmt.Println("vector2:", vector2)

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			matrix[i][j] = i + j
		}
	}
	fmt.Println("matrix:", matrix)

}

func getAverage(array []int, size int) float32 {
	var sum int
	var average float32

	for i := 0; i < size; i++ {
		sum += array[i]
	}
	average = float32(sum / size)
	return average
}

func main() {
	simpleArrays()
	var balance = []int{1, 2, 3, 4, 5}
	fmt.Println("avg:", getAverage(balance, 5))
}
