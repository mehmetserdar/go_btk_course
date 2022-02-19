package variables

import "fmt"

func Variables() {

	metin := "Merhaba Türkiye!"
	fmt.Print("Merhaba Dünya! \n")
	fmt.Println(metin)

	//integer
	var kdv int = 10

	fmt.Println(kdv)
	fmt.Println(100 + 100*kdv/100)

	//float
	var kdv2 float32 = .15
	fmt.Println(kdv2)
	fmt.Println(100 + 100*kdv2)

	kdv3 := .08

	fmt.Println(kdv3)
	fmt.Printf("Veri tipi: %T\n", kdv3)

	var durum bool
	metin1 := "Mehmet"
	metin2 := "Serdar"

	durum = metin1 != metin2
	fmt.Println(durum)
	fmt.Println(!durum)

}
