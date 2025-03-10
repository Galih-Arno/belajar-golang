package main
import "fmt"

// deklarasi beberapa varible dalam satu baris sekaligus yang sama
// func main()  {
// 	// var a, b, c, d int = 1, 3, 5, 8

// 	// fmt.Println(a)
// 	// fmt.Println(b)
// 	// fmt.Println(c)
// 	// fmt.Println(d)

// 	var a, b = 5, "Hello"
// 	c, d := 3, "World!"

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// }

// deklarasi variable go dalam blok
func main()  {
	var (
		a string = "Hello"
		b int = 2
		c = "World!"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}