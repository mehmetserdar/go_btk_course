package defer_statement

import "fmt"

func Demo1(sayi int32) string {
	defer DeferFunc()
	if sayi%2 == 0 {
		return "çift sayıdır"
	} else {
		return "tek sayıdır"
	}

}

func Test() {

	sonuc := Demo1(11)
	fmt.Println(sonuc)

}

func DeferFunc() {
	fmt.Println("bitti")
}
