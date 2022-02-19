package forrange

import "fmt"

//1-10 arasındaki tek sayıları topla
//for-range
func Demo1() {
	sayilar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	toplam := 0

	for _, sayi := range sayilar {
		if sayi%2 != 0 {
			toplam += sayi
		}
	}
	fmt.Println(toplam)
}
