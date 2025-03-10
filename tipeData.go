package main
import "fmt"

// Tipe data adalah konsep penting dalam pemrograman. Tipe data menentukan ukuran dan jenis nilai variabel.
// Go diketik secara statis, artinya setelah jenis variabel ditentukan, ia hanya 
// dapat menyimpan data dari jenis tersebut.

// Go memiliki tiga tipe data dasar:
// bool: mewakili nilai boolean dan benar atau salah
// Numerik: mewakili jenis bilangan bulat, nilai floating point, dan jenis kompleks
// string: mewakili nilai string

// contoh
// func main()  {
// 	var(
// 		a bool = true				// boolean
// 		b int = 5					// integer
// 		c float32 = 3.14			// floating point number
// 		d string = "John Doe"		// string
// 	)

// 	fmt.Println("Boolean: ", a)
// 	fmt.Println("Integer: ", b)
// 	fmt.Println("Float: ", c)
// 	fmt.Println("String: ", d)
// }

// ========================================================================

// Tipe Data Boolean
// Tipe data boolean dideklarasikan dengan bool kata kunci dan hanya dapat mengambil nilai true atau false.
// Nilai default tipe data boolean adalah false.

// contoh
// func main()  {
// 	var (
// 		a bool = true	// tipe deklarasi dengan nilai awal
// 		b = true		// tanpa tipe deklarasi dengan nilai awal
// 		c bool			// tipe deklarasi tanpa nilai awal
// 	)
// 	d := true			// tanpa tipe deklarasi dengan nilai awal

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// }

// =======================================================================================

// Tipe Data Integer
// Tipe data bilangan bulat digunakan untuk menyimpan bilangan bulat tanpa desimal, seperti 35, -50, atau 1345000.
// Tipe data bilangan bulat memiliki dua kategori:
// Bilangan bulat bertanda - dapat menyimpan nilai positif dan negatif
// Bilangan bulat yang tidak ditandatangani - hanya dapat menyimpan nilai non-negatif
// Jenis default untuk bilangan bulat adalah int. Jika Anda tidak menentukan jenis, jenisnya akan menjadi int.

// contoh
// Bilangan bulat bertandatangan, dideklarasikan dengan salah satu int kata kunci, dapat menyimpan nilai positif dan negatif:
// func main()  {
// 	var (
// 		x int = 500
// 		y int = -2100
// 	)

// 	fmt.Printf("Tipe: (%T), nilainya: (%v)\n", x, x)
// 	fmt.Printf("Tipe: (%T), nilainya: (%v)", y, y)

// 	// Go memiliki lima kata kunci/jenis bilangan bulat yang ditandatangani:
// 	// int 		Depends on platform:
// 	// 			32 bits in 32 bit systems and		-2147483648 to 2147483647 in 32 bit systems and
// 	// 			64 bit in 64 bit systems 			-9223372036854775808 to 9223372036854775807 in 64 bit systems
// 	// int8 	8 bits/1 byte 						-128 to 127
// 	// int16 	16 bits/2 byte 						-32768 to 32767
// 	// int32 	32 bits/4 byte 						-2147483648 to 2147483647
// 	// int64 	64 bits/8 byte 						-9223372036854775808 to 9223372036854775807
// }

// bilangan bulat tanpa tanda, dideklarasikan dengan salah satu uint kata kunci, hanya dapat menyimpan nilai non-negatif:
// func main()  {
// 	var (
// 		x uint = 300
// 		y uint = 21002
// 	)

// 	fmt.Printf("Tipe: (%T), nilainya: (%v)\n", x, x)
// 	fmt.Printf("Tipe: (%T), nilainya: (%v)", y, y)

// 	// Go memiliki lima kata kunci/jenis bilangan bulat yang tidak ditandatangani:
// 	// uint 	Depends on platform:
// 	// 			32 bits in 32 bit systems and		0 to 4294967295 in 32 bit systems and
// 	// 			64 bit in 64 bit systems 			0 to 18446744073709551615 in 64 bit systems
// 	// uint8 	8 bits/1 byte 						0 to 255
// 	// uint16 	16 bits/2 byte 						0 to 65535
// 	// uint32 	32 bits/4 byte 						0 to 4294967295
// 	// uint64 	64 bits/8 byte 						0 to 18446744073709551615
// }


// Jenis bilangan bulat mana yang akan digunakan?
// Jenis bilangan bulat yang akan dipilih, tergantung pada nilai yang harus disimpan variabel.

// Contoh ini akan menghasilkan kesalahan karena 1000 berada 
// di luar jangkauan untuk int8 (yaitu dari -128 hingga 127):
// func main()  {
// 	var x int8 = 20000
// 	fmt.Printf("Tipe: (%T) nilai: %v", x, x)
// 	// ./tipeData.go:104:15: cannot use 20000 (untyped int constant) as int8 value in variable declaration (overflows)
// }


// =================================================================================================


// Tipe Data Float
// Tipe data float digunakan untuk menyimpan angka positif dan negatif dengan titik desimal, seperti 35.3, -2.34, atau 3597.34987.
// Jenis data float memiliki dua kata kunci:
// float32 	32 bits 	-3.4e+38 to 3.4e+38.
// float64 	64 bits 	-1.7e+308 to +1.7e+308.
// Jenis default untuk float adalah float64. Jika Anda tidak menentukan jenis, jenisnya akan menjadi float64.

// contoh
// func main()  {
// 	var (
// 		x float32 = 123.12
// 		y float32 = 3.4e+38
// 		z float64 = 1.7e+308
// 	)

// 	fmt.Printf("Tipe (%T) dengan nilai %v\n", x, x)
// 	fmt.Printf("Tipe (%T) dengan nilai %v\n", y, y)
// 	fmt.Printf("Tipe (%T) dengan nilai %v\n", z, z)
// }

// Jenis float mana yang akan digunakan?
// Jenis float yang akan dipilih, tergantung pada nilai yang harus disimpan variabel.

// contoh salah
// func main()  {
// 	var x float32 = 3.4e+39
// 	fmt.Printf(x)
// 	// # command-line-arguments
// 	// ./tipeData.go:140:18: cannot use 3.4e+39 (untyped float constant) as float32 value in variable declaration (overflows)
// 	// ./tipeData.go:141:13: cannot use x (variable of type float32) as string value in argument to fmt.Printf
// }


// ===========================================================================================


// Tipe Data String
// Tipe data string digunakan untuk menyimpan urutan karakter (teks). 
// Nilai string harus dikelilingi oleh tanda kutip ganda:

// contoh
func main()  {
	var (
		txt1 string = "Hello!"
		txt2 string
	)
	txt3 := "John Doe"

	fmt.Printf("Tipe (%T) dengan hasil %v\n", txt1, txt1)
	fmt.Printf("Tipe (%T) dengan hasil %v\n", txt2, txt2)
	fmt.Printf("Tipe (%T) dengan hasil %v\n", txt3, txt3)
}