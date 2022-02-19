package pointer

import "fmt"

func Pointer(sayi *int) {
	*sayi++
	fmt.Println("Demodaki sayı:", *sayi)
}
