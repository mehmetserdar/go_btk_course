package pointer

import "fmt"

func Demo1(sayilar []int) {
	sayilar[0] = 100
	fmt.Println("Demodaki sayı:", sayilar[0])
}
