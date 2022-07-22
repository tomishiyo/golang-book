package main

import "fmt"

func main() {
	test_fib()
}

/* Write a function that takes an integer and halves it and return true if it
was even or false if it was odd. For example half(1) should return (0, false)
and half(2) should return (1, true) */

func half(number int) (quotient float64, status_code bool) {
	quotient = float64(number) / 2

	if number%2 == 0 {
		return quotient, true
	} else {
		return quotient, false
	}
}

func test_half() {
	var (
		quotient    float64
		status_code bool
	)

	for i := 1; i <= 50; i++ {
		quotient, status_code = half(i)
		fmt.Println(quotient, status_code)
	}
}

/*Write a function with one variadic parameter that finds the greatest number
in a list of numbers*/

func find_greatest(numbers ...int) (greatest_number int) {
	greatest_number = numbers[0]

	for _, number := range numbers {
		if number > greatest_number {
			greatest_number = number
		}
	}

	return
}

/* Using makeEvenGenerator as an example, write a makeOddGenerator function that
generates odd numbers */

func makeOddGenerator() func() uint {
	counter := uint(1)

	return func() (odd uint) {
		odd = counter
		counter += 2
		return
	}
}

func testOddGenerator() {
	nextOdd := makeOddGenerator()

	for i := 0; i <= 10; i++ {
		fmt.Println(nextOdd())
	}

}

/* The Fibonnaci sequence is defined as: fib(0)=0, fib(1)=1,
fib(n) = fib(n-1) + fib(n-2). Write a recursive function which can find fib(n).*/

func fib(n int) (an int) {
	if n == 0 {
		an = 0
	} else if n == 1 {
		an = 1
	} else {
		an = fib(n-1) + fib(n-2)
	}
	return
}

func test_fib() {
	for i := 0; i <= 10; i++ {
		fmt.Println(fib(i))
	}
}
