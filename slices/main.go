package main

import "fmt"

func main() { //yeni eleman geldiğinde eklemek arraylare göre çok daha kolay
	isimler := make([]string, 3)

	isimler[0] = "ferhat"
	isimler[1] = "altay"
	isimler[2] = "szlai"
	isimler = append(isimler, "pelkas") //ek olrak bir isim daha girdik
	fmt.Println(isimler)
}
