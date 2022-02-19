package conditionals

import "fmt"

func Ws1() {
	// 3 adet int değişken . ekrana en büyük olanı yazdırın
	a := 10
	b := 120
	c := 3400

	// if a > b || a > c {
	// 	fmt.Print(a)
	// } else if b > a || b > c {
	// 	fmt.Print(b)
	// } else {
	// 	fmt.Print(c)
	// }

	enBuyuk := a

	if b > enBuyuk {
		enBuyuk = b
	}

	if c > enBuyuk {
		enBuyuk = c
	}

	fmt.Printf("En büyük sayı: %v", enBuyuk)
}
