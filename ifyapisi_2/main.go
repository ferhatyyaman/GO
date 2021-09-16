package main

import "fmt"

func main() { // 3 sayi arasindan en buyugunu bulma
	var sayi1, sayi2, sayi3 int = 9, 6, 7

	if sayi1 < sayi2 {
		if sayi2 > sayi3 {
			fmt.Println("en büyük sayi sayi2 dir")
		} else {
			fmt.Println("en büyük sayi sayi3 dir")
		}
	} else if sayi1 > sayi2 {

		if sayi1 > sayi3 {
			fmt.Println("en büyük sayi sayi1 dir")
		} else {
			fmt.Println("en büyük sayi sayi3 dir")
		}
	}
}
