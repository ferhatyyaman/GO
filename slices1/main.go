package main

import "fmt"

func main() {
	sehirler := []string{"ankara", "istanbul", "izmir"}
	fmt.Println(sehirler)
	sehirlerKopya := make([]string, len(sehirler)) //sehirlerin eleman sayisi kadar olusturduk
	fmt.Println(sehirlerKopya)
	copy(sehirlerKopya, sehirler) //sehirlerkopyaya sehirlerin eleman sayisini atadik
	fmt.Println(sehirlerKopya)

	sehirler = append(sehirler, "antalya") //sehirlere yeni eleman eklersek
	fmt.Println(sehirler)
	fmt.Println(sehirlerKopya) //eleman sayisinda değişiklik olmaz

	fmt.Println(sehirler[1:3]) //1. indisten basla 3. indisi dahil etme

}
