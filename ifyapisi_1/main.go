package main

import "fmt"

func main() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("hesaptaki para yetersiz")
	}
	if cekilmekIstenen < hesap {
		fmt.Println("paranız hazırlanıyor ")
		hesap = hesap - cekilmekIstenen

	}
	fmt.Printf("hesapta kalan para: %v", hesap)
	//-----------2. yöntem---------
	//fmt.Println("hesapta kalan para : "+fmt.sprintf(%v,hesap))
}
