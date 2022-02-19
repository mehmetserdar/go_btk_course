package slices

import "fmt"

func Slices() {
	isimler := make([]string, 3)
	fmt.Println(isimler)
	isimler[0] = "Mehmet"
	isimler[1] = "Ahmet"
	isimler[2] = "Ayşe"
	isimler = append(isimler, "Fatma")
	fmt.Println(isimler)
	fmt.Println(len(isimler))
}
