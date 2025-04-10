package main
import "fmt"

// Go menawarkan beberapa kata kerja pemformatan yang dapat digunakan dengan fungsi Printf()

// pemformatan kata kerja yang umum
// %v mencetak nilai secara default
// %#v mencetak nilai dengan format yang digunakan dalam Go-sintaks
// %T mencetak tipe nilai
// %% mencetak simbol %

// contoh
// func main()  {
// 	var i = 12.4
// 	var txt = "Hello World!"

// 	fmt.Printf("%v\n", i)
// 	fmt.Printf("%#v\n", i)
// 	fmt.Printf("%v%%\n", i)
// 	fmt.Printf("%T\n", i)

// 	fmt.Printf("%v\n", txt)
// 	fmt.Printf("%#v\n", txt)
// 	fmt.Printf("%T\n", txt)
// }


// Kata kerja berikut dapat digunakan dengan tipe data bilangan bulat:
// %b 		Base 2
// %d 		Base 10
// %+d 		Base 10 and always show sign
// %o 		Base 8
// %O 		Base 8, with leading 0o
// %x 		Base 16, lowercase
// %X 		Base 16, uppercase
// %#x 		Base 16, with leading 0x
// %4d 		Pad with spaces (width 4, right justified)
// %-4d 	Pad with spaces (width 4, left justified)
// %04d 	Pad with zeroes (width 4

// contoh
// func main() {
// 	var i = 15

// 	fmt.Printf("%b\n", i)
// 	fmt.Printf("%d\n", i)
// 	fmt.Printf("%+d\n", i)
// 	fmt.Printf("%o\n", i)
// 	fmt.Printf("%O\n", i)
// 	fmt.Printf("%x\n", i)
// 	fmt.Printf("%X\n", i)
// 	fmt.Printf("%#x\n", i)
// 	fmt.Printf("%4d\n", i)
// 	fmt.Printf("%-4d\n", i)
// 	fmt.Printf("%04d\n", i)
// }


// Kata kerja berikut dapat digunakan dengan tipe data string:
// %s 		Prints the value as plain string
// %q 		Prints the value as a double-quoted string
// %8s 		Prints the value as plain string (width 8, right justified)
// %-8s 	Prints the value as plain string (width 8, left justified)
// %x 		Prints the value as hex dump of byte values
// % x 		Prints the value as hex dump with spaces

// contoh
// func main() {
// 	var txt = "Hello"

// 	fmt.Printf("%s\n", txt)
// 	fmt.Printf("%q\n", txt)
// 	fmt.Printf("%8s\n", txt)
// 	fmt.Printf("%-8s\n", txt)
// 	fmt.Printf("%x\n", txt)
// 	fmt.Printf("% x\n", txt)
// }


// Kata kerja berikut dapat digunakan dengan tipe data boolean:
// %t 	Value of the boolean operator in true or false format (same as using %v)

// contoh
// func main() {
// 	var i = true
// 	var j = false

// 	fmt.Printf("%t\n", i)
// 	fmt.Printf("%t\n", j)
// }


// Kata kerja berikut dapat digunakan dengan tipe data float: 
// %e 		Scientific notation with 'e' as exponent
// %f 		Decimal point, no exponent
// %.2f 	Default width, precision 2
// %6.2f 	Width 6, precision 2
// %g 		Exponent as needed, only necessary digits

// contoh
func main() {
	var i = 3.141

	fmt.Printf("%e\n", i)
	fmt.Printf("%f\n", i)
	fmt.Printf("%.2f\n", i)
	fmt.Printf("%6.2f\n", i)
	fmt.Printf("%g\n", i)
}