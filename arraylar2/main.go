package main

import "fmt"

func main() {

	sayilar := [5]int{20, 30, 50, 10, 2}

	enbuyuk := sayilar[0]

	for i := 0; i < len(sayilar); i++ {
		if sayilar[i] > enbuyuk {
			enbuyuk = sayilar[i]
		}
	}
	fmt.Println("en büyük : ", enbuyuk)
}
