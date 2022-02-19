package main

import (
	"fmt"
	"gobtk/project"
)

func main() {
	fmt.Println("GO KURSU BTK")
	//variables.Variables()
	//conditionals.Ws1()
	//functions.SelamVer("Serdar")
	//var sonuc = functions.Topla(2, 7)
	//fmt.Println(sonuc)

	// sonuc1, sonuc2, sonuc3, _ := functions.DortIslem(10, 5)

	// fmt.Println("Toplam: ", sonuc1)
	// fmt.Println("Fark: ", sonuc2)
	// fmt.Println("Çarpım: ", sonuc3)
	//fmt.Println("Bölüm: ", sonuc4)
	// var sonuc = functions.ToplaVariadic(2, 45, 76, 87, 12)
	// fmt.Println(sonuc)

	// sayilar := []int{3, 5, 7, 6, 8}
	// fmt.Println(functions.ToplaVariadic(sayilar...))

	// number := 20
	// pointer.Pointer(&number)
	// fmt.Println("Maindeki sayı:", number)

	// sayilar := []int{1, 2, 3}
	// pointer.Demo1(sayilar)
	// fmt.Println("Maindeki sayı:", sayilar[0])

	// ciftSayiToplamC := make(chan int)
	// tekSayiToplamC := make(chan int)
	// go channels.CiftSayilar(ciftSayiToplamC)
	// go channels.TekSayilar(tekSayiToplamC)

	// ciftSayiToplam, tekSayiToplam := <-ciftSayiToplamC, <-tekSayiToplamC

	// carpim := ciftSayiToplam * tekSayiToplam
	// fmt.Println("Çarpım:", carpim)

	// fmt.Println(error_handling.TahminEt2(99))

	//project.AddProduct()
	product, _ := project.AddProduct()
	fmt.Println(product)

	products, _ := project.GetAllProducts()

	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}
}
