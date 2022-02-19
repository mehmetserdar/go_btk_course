package string_functions

import (
	"fmt"
	s "strings"
)

//case sensitive
//ascii
func Demo1() {
	isim := "Serdar"
	fmt.Println(s.Count(isim, "r"))
	fmt.Println(s.Contains(isim, "r"))
	sonuc := s.Index(isim, "e")

	if sonuc != -1 {
		fmt.Println("e harfi var")
	} else {
		fmt.Println("e harfi yok")
	}
	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))

}
