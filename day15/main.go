package main

import "fmt"

func main() {
	a := 883
	b := 879

	count := 0

	for i := 0; i < 40*1000*1000; i++ {
		a = (a * 16807) % 2147483647
		b = (b * 48271) % 2147483647

		if (a & 0xFFFF) == (b & 0xFFFF) {
			count++
		}
	}

	fmt.Println(count)
}
