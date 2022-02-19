package maps

import "fmt"

func Maps() {
	// key-value
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["pencil"] = "Kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("Eleman Sayisi: ", len(sozluk))
	fmt.Println("Sözlük: ", sozluk)
	delete(sozluk, "book")
	fmt.Println("Eleman Sayisi: ", len(sozluk))
	fmt.Println("Sözlük: ", sozluk)

	deger, varMi := sozluk["book"]
	fmt.Println(deger)
	fmt.Println("Listede olma durumu:", varMi)

	sozluk2 := map[string]string{"glass": "bardak", "fire": "ateş"}
	fmt.Println(sozluk2)

}
