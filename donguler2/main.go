package main

import "fmt"

func main() {

	tutugumSayi := 50
	tahtahminEdilenSayi := 0

	fmt.Println("Bir sayi tutunuz")
	fmt.Scanln(&tahtahminEdilenSayi)

	for tutugumSayi != tahtahminEdilenSayi {
		if tutugumSayi < tahtahminEdilenSayi {
			fmt.Printf("büyük sayi girmelisin %v ", tahtahminEdilenSayi-tutugumSayi)
			fmt.Println("az gir")
			fmt.Scanln(&tahtahminEdilenSayi)
		} else if tutugumSayi > tahtahminEdilenSayi {
			fmt.Printf("kücük sayi girmelisin %v ", tutugumSayi-tahtahminEdilenSayi)
			fmt.Println("fazla gir")
			fmt.Scanln(&tahtahminEdilenSayi)
		} else if tutugumSayi == tahtahminEdilenSayi {

		}

	}
	fmt.Println("bravo bildin")
}
