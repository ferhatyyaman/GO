package main

import "fmt"

func main() {
	sayac := 0
	sayi := 0
	sayi2 := 0
	sayac2 := 0

	fmt.Println("ilk sayiyi yazın")
	fmt.Scanln(&sayi)
	for i := 1; i < sayi; i++ {

		if sayi%i == 0 {
			sayac = sayac + i
		}
	}
	fmt.Println("ilk sayiyi yazın")
	fmt.Scanln(&sayi2)
	for i := 1; i < sayi2; i++ {

		if sayi2%i == 0 {
			sayac2 = sayac2 + i
		}
	}

	if sayac == sayi2 && sayac2 == sayi {
		fmt.Println("arkadaş sayılar")
	} else {
		fmt.Println("arkadas sayı degil")
	}

}
