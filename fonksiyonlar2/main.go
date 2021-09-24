package main

import "fmt"

func dortİslem(sayi1 int, sayi2 int) (int, int, int, float32) {
	toplam := sayi1 + sayi2
	cikarim := sayi1 - sayi2
	carpim := sayi1 * sayi2
	bolum := float32(sayi1 / sayi2)

	return toplam, cikarim, carpim, bolum

}

func main() {

	sonuc1, sonuc2, sonuc3, sonuc4 := dortİslem(10, 6)
	fmt.Println("topolam= ", sonuc1)
	fmt.Println("topolam= ", sonuc2)
	fmt.Println("topolam= ", sonuc3)
	fmt.Println("topolam= ", sonuc4)

}
