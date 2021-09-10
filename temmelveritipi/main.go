package main

import "fmt"

func main() {

	var name = "ferhat"
	var age int16 = -256
	var isMarried bool = true
	var weight float32 = 72.5

	isMarried = false

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMarried) //veri tipini yazdırır
	fmt.Printf("%T\n", weight)

}
