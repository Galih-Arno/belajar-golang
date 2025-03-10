package main
import "fmt"

// const digunakan untuk mendeklarasikan variable sebagai konstanta yang berarti nilai tidak dapat diubah
// konstanta biasanya ditulis dengan huruf kapital untuk memudahkan identifikasi dan membedakan dari variable
// konstanta dapat dideklarasikan didalam maupun diluar fungsi.
// const PI = 3.14

// func main()  {
// 	fmt.Println(PI)
// }

// const A uint8 = 1
// func main()  {
// 	fmt.Println(A)
// }

const (
	A int = 1
	B = 3.14
	C = "Hai!"
)

func main()  {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}