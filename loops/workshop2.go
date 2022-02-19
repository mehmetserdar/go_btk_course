package loops

import "fmt"

// Kullanıcıdan bir sayı girmesini isteyin.
// 23: Asaldır.
// Sayının asal olup olmadığını bulan program.
func Ws2() {
	sayi := 0

	fmt.Println("ASAL MI DEĞİL Mİ?")
	fmt.Println("Lütfen bir sayı giriniz")
	fmt.Scanln(&sayi)

	asalMi := true
	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			asalMi = false
		}
	}

	if asalMi {
		fmt.Printf("%v: Asaldır", sayi)
	} else {
		fmt.Printf("%v: Asal değildir", sayi)
	}

}
