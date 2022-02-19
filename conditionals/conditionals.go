package conditionals

import "fmt"

func Conditionals() {
	var hesap float64 = 3900
	var cekilmekIstenen float64 = 1000

	if cekilmekIstenen > hesap {
		fmt.Printf("Hesaptaki para yetersiz. Hesabınız bakiyesi: %v", hesap)
	} else if cekilmekIstenen == hesap {
		fmt.Print("Paranız hazırlanıyor")
		fmt.Print("Hesapta para kalmadı.")
	} else {
		fmt.Print("Paranız hazırlanıyor.")
	}

}
