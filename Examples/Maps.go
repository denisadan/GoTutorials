package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("map length:", len(m))
	delete(m, "k2")
	fmt.Println("map:", m)
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	presAge := make(map[string]int)
	presAge["Roosvelt"] = 42
	presAge["Kennedy"] = 43
	fmt.Println("second map length:", len(presAge))
	fmt.Println("second map:", presAge)
	delete(presAge, "Roosvelt")
	fmt.Println("second map after deletion:", presAge)

	//map inside a map

}

