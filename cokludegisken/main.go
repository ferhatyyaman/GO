package main

import "fmt"

func main() {

	/*	-------------1.yöntem-----------------
		 	var (
			name      string = "ferhat"
			age       int    = 22
			isMarried bool   = true
			weight float32 = 72.5
			height int     = 178
		) */

	/* --------------2.yöntem------------------
	var name, age, isMarried, weight, height = "ferhat", 22, true, 72.5, 178 */

	/* ------------3.yöntem-------------------*/
	name, age, isMarried, weight, height := "ferhat", 22, true, 72.5, 178

	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)
	fmt.Println(height)

}
