package loops

import (
	"fmt"
	"math/rand"
	"time"
)

//
func Ws1() {
	// 0-100 arasında random değişen sayı verir
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 100
	aklimdakiSayi := rand.Intn(max-min+1) + min

	tahminSayi := 0
	i := 0
	var sonuc string = ""

	fmt.Println("Lütfen 0 ve 100 Arasında Bir Sayı Tahmin Ediniz")
	fmt.Scanln(&tahminSayi)
	i++

	if tahminSayi > 0 && tahminSayi <= 100 {

		for aklimdakiSayi != tahminSayi {
			if aklimdakiSayi > tahminSayi {
				fmt.Println("Daha büyük sayı giriniz")
				fmt.Scanln(&tahminSayi)
				i++
			}

			if aklimdakiSayi < tahminSayi {
				fmt.Println("Daha küçük sayı giriniz")
				fmt.Scanln(&tahminSayi)
				i++
			}
		}

		if i <= 3 {
			sonuc = "Süper"
		} else if i <= 6 {
			sonuc = "Güzel"
		} else {
			sonuc = "Fena değil"
		}

		fmt.Printf("%v, %v Tahminde Bildiniz", sonuc, i)

	} else {
		println("Giridğiniz sayı 0 ila 100 arasında olmalıdır.")
	}

}
