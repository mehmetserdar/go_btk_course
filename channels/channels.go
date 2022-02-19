package channels

func CiftSayilar(c chan int) {
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		toplam = toplam + i
	}

	c <- toplam
}
func TekSayilar(c chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam = toplam + i
	}

	c <- toplam
}
