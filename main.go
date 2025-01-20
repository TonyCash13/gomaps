package main

import "fmt"

func main() {

	mapsMe := map[int]string{

		1: "Pavel",
		2: "Sergey",
		3: "Alexander",
	}
	count := 0
	for key, _ := range mapsMe {
		count++
		fmt.Println(key)
	}
	fmt.Println("Количество ключей", count)
}
