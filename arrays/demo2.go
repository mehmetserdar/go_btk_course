package arrays

import "fmt"

func Demo2() {
	sayilar := [5]int{20, 22, 39, 54, 98}
	fmt.Println(sayilar)

	enBuyuk := sayilar[0]

	for i := 0; i < len(sayilar); i++ {
		if sayilar[i] > enBuyuk {
			enBuyuk = sayilar[i]
		}

	}
	fmt.Println("En büyük: ", enBuyuk)

}
