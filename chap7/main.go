/* Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y))
should give you x=2 and y=1 */
package main

import "fmt"

func swap(x *int, y *int) {
	tmp_x := *x
	*x = *y
	*y = tmp_x
}

func main() {
	x := 1
	y := 2

	fmt.Printf("Pre swap: x= %v, y= %v \n", x, y)

	swap(&x, &y)

	fmt.Printf("After swap: x= %v, y= %v \n", x, y)
}
