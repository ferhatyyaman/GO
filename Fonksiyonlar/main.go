package main

import "fmt"

func selamVer(kullaniciadi string) {
	fmt.Println("merhaba", kullaniciadi)
}

func topla(sayi1 int, sayi2 int) int {
	var toplam int = sayi1 + sayi2
	return toplam

}
func main() {

	var toplam = topla(2, 6)
	fmt.Println(toplam)
	selamVer("ferhat")

}
