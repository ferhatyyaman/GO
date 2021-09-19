package main

import "fmt"

func main() {

	var sehirler [5]string
	sehirler[0] = "istanbul"
	sehirler[1] = "ankara"
	sehirler[3] = "izmir"
	sehirler[4] = "antalya"
	fmt.Println(sehirler)

	for i := 0; i < 5; i++ {
		fmt.Println(sehirler[i])
	}
}
