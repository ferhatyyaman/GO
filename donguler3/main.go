package main

import "fmt"

func main() {

	tutugumSayi := 50
	tahtahminEdilenSayi := 0
	sayac := 0

	fmt.Println("Bir sayi tutunuz")
	fmt.Scanln(&tahtahminEdilenSayi)

	for tutugumSayi != tahtahminEdilenSayi {
		if tutugumSayi < tahtahminEdilenSayi {
			fmt.Printf("büyük sayi girmelisin %v ", tahtahminEdilenSayi-tutugumSayi)
			fmt.Println("az gir")
			fmt.Scanln(&tahtahminEdilenSayi)
			sayac++

		} else if tutugumSayi > tahtahminEdilenSayi {
			fmt.Printf("kücük sayi girmelisin %v ", tutugumSayi-tahtahminEdilenSayi)
			fmt.Println("fazla gir")
			fmt.Scanln(&tahtahminEdilenSayi)
			sayac++
		} else if tutugumSayi == tahtahminEdilenSayi {

		}

	}
	fmt.Printf(" deneme sayisi %v", sayac)
	if 0 < sayac && sayac <= 3 {
		fmt.Println(" super")
	} else if 3 < sayac && sayac <= 6 {
		fmt.Println(" guzel")
	} else if 6 < sayac {
		fmt.Println(" fena degil")
	}

}
