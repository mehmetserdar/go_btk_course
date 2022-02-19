package structs

import "fmt"

func Structs() {
	fmt.Println(product{"Bilgisayar", 500, "XYZ"})
	fmt.Println(product{"Televizyon", 300, "ABC"})
	fmt.Println(product{name: "Bilgisayar"})

}

type product struct {
	name      string
	unitPrice float64
	brand     string
}
