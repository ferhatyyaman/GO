package main

import "fmt"

func main() {

	var sayilar [2][3]int

	sayilar[0][0] = 1
	sayilar[0][1] = 3
	sayilar[0][2] = 5
	sayilar[1][0] = 7
	sayilar[1][1] = 9
	sayilar[1][2] = 11

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {

			fmt.Print(sayilar[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}

}
