package main
import "fmt"

/*
	Go memiliki 3 fungsi untuk menampilkan teks
	- Print()
	- Println()
	- Printf()
*/

// Print()

// func main()  {
// 	var i, j string = "Hello", "World!"

// 	fmt.Print(i)
// 	fmt.Print(j)

// 	mencetak argumen dalam baris baru
// 	fmt.Print(i, "\n")
// 	fmt.Print(j, "\n")

// 	bisa juga menggunakan satu Print() untuk mencetak beberapa variable
// 	fmt.Print(i, "\n", j)

// 	jika ingin menambahkan spasi
// 	fmt.Print(i, " ", j)

// 	jika tidak ada ada string, Print() menyisipkan spasi di antara argumen
// 	var i, j = 10, 12

// 	fmt.Print(i, j)
// }


// Println()

// fungsi dari Println() mirip Print() dengan perbedaan bahwa spasi ditambahkan di antara arguman,
// dan baris baru di tambahkan di akhir

// func main()  {
// 	var i, j string = "Hello", "World"

// 	fmt.Println(i, j)
// }


// Printf()

// fungsi Printf() memformat argumennya berdasarkan kata kerja pemformatan yang 
// diberikan dan kemudian mencetaknya
// misalnya %v digunakan untuk mencetak nilai dari argumen,
// 			%T digunakan untuk mencetak jenis dari argumen

func main()  {
	var i string = "Hello World"
	var j int = 12

	fmt.Printf("nilai dari i: %v dan tipenya adalah (%T)\n", i, i)
	fmt.Printf("nilai dari j: %v dan tipenya adalah (%T)\n", j, j)
}
