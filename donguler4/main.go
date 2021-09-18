package main

import "fmt"

func main() {
	sayi := 0
	asalmi := true
	fmt.Println("Bir sayi tutunuz")
	fmt.Scanln(&sayi)

	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			asalmi = false
		}
	}
	if asalmi == true {
		fmt.Println("asaldır")
	} else {
		fmt.Println("asal degildir")
	}

}
