package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Mehmet"
	fmt.Println(s.HasPrefix(isim, "Meh"))
	fmt.Println(s.HasSuffix(isim, "et"))
	fmt.Println(s.Index(isim, "et"))
	harfler := []string{"m", "e", "h", "m", "e", "t"}
	sonuc := s.Join(harfler, "*")
	fmt.Println(sonuc)

	fmt.Println(s.Replace(sonuc, "*", "+", 3))
	fmt.Println(s.Replace(sonuc, "*", "+", -1))

	fmt.Println(s.Split(sonuc, "*"))
	fmt.Println(s.Repeat(sonuc, 5))
}
