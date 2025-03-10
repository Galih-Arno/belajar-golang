package main
import "fmt"

// Slice mirip dengan array, tetapi lebih kuat dan fleksibel.
// Seperti array, irisan juga digunakan untuk menyimpan beberapa nilai dari jenis yang sama dalam satu variabel.
// Namun, tidak seperti array, panjang irisan dapat tumbuh dan menyusut sesuai keinginan Anda.
// Di Go, ada beberapa cara untuk membuat irisan:
// Menggunakan format []datatype{values}
// Membuat irisan dari array
// Menggunakan fungsi make()

// Membuat Irisan Dengan []datatype{values}
// contoh
// func main()  {
// 	slice1 := [] int {}
// 	fmt.Println(len(slice1))	// len() fungsi - mengembalikan panjang irisan (jumlah elemen dalam irisan)
// 	fmt.Println(cap(slice1))	// cap() function - mengembalikan kapasitas irisan (jumlah elemen yang dapat tumbuh atau menyusut ke irisan)
// 	fmt.Println(slice1, "\n")

// 	slice2 := [] string {"Go", "Slices", "Are", "Powerful"}
// 	fmt.Println(len(slice2))
// 	fmt.Println(cap(slice2))
// 	fmt.Println(slice2)
// }

// ===================================================================================

// Membuat Irisan Dari Array
// Contoh ini menunjukkan cara membuat irisan dari array:
// func main()  {
// 	arr1 := [6] int {10,11,12,14,19,21}
// 	slice1 := arr1[2:4]

// 	fmt.Printf("my slice = %v\n", slice1)
// 	fmt.Printf("Panjangnya = %d\n", len(slice1))
// 	fmt.Printf("kapasitas = %d\n", cap(slice1))
// }

// Pada contoh di atas slice1 adalah irisan dengan panjang 2. 
// Itu terbuat dari arr1 yang merupakan array dengan panjang 6.

// Irisan dimulai dari elemen ketiga dari array yang memiliki nilai 12 (ingat bahwa indeks array dimulai dari 0. 
// Itu berarti bahwa [0] adalah elemen pertama, [1] adalah elemen kedua, dll.). 
// Irisan dapat tumbuh ke akhir array. Artinya, kapasitas irisan adalah 4.
// Jika slice1 dimulai dari elemen 0, kapasitas irisan akan menjadi 6.


// =====================================================================================


// Buat Slice Dengan Fungsi make()
// Fungsi make() juga dapat digunakan untuk membuat irisan.
// Jika parameter kapasitas tidak ditentukan, itu akan sama dengan panjang.
// Contoh ini menunjukkan cara membuat irisan menggunakan make() fungsi:
// func main()  {
// 	myslice1 := make([] int, 5, 10)
// 	fmt.Printf("my slice 1 = %v\n", myslice1)
// 	fmt.Printf("panjang = %d\n", len(myslice1))
// 	fmt.Printf("kapasitas = %d\n", cap(myslice1))

// 	myslice2 := make([] int, 5)
// 	fmt.Printf("my slice 2 = %v\n", myslice2)
// 	fmt.Printf("panjang = %d\n", len(myslice2))
// 	fmt.Printf("kapasitas = %d\n", cap(myslice2))
// }


// =================================================================================================


// Akses, Ubah, Tambahkan dan Salin Irisan

// Akses elemen irisan
// Contoh ini menunjukkan cara mengakses elemen pertama dan ketiga dalam potongan harga:
// func main() {
// 	prices := []int{10,20,30}

// 	fmt.Println(prices[0])
// 	fmt.Println(prices[2])
// }


// ------------------------------------------------------------------------------------------------


// Ubah elemen irisan
// Contoh ini menunjukkan cara mengubah elemen ketiga dalam potongan harga:
// func main() {
// 	prices := []int{10,20,30}
// 	prices[2] = 50
// 	fmt.Println(prices[0])
// 	fmt.Println(prices[2])
// }


// ------------------------------------------------------------------------------------------------


// Tambahkan Elemen Ke Irisan
// Anda dapat menambahkan elemen ke akhir slice menggunakan fungsi append()
// Contoh ini menunjukkan cara menambahkan elemen ke akhir irisan:
// func main() {
// 	myslice1 := []int{1, 2, 3, 4, 5, 6}
// 	fmt.Printf("myslice1 = %v\n", myslice1)
// 	fmt.Printf("length = %d\n", len(myslice1))
// 	fmt.Printf("capacity = %d\n\n", cap(myslice1))

// 	myslice1 = append(myslice1, 20, 21)
// 	fmt.Printf("myslice1 = %v\n", myslice1)
// 	fmt.Printf("length = %d\n", len(myslice1))
// 	fmt.Printf("capacity = %d\n", cap(myslice1))
// }


// -----------------------------------------------------------------------------------------------


// Tambahkan Satu Irisan Ke Irisan Lain
// Untuk menambahkan semua elemen dari satu irisan ke irisan lain, gunakan fungsi append()
// Contoh ini menunjukkan cara menambahkan satu irisan ke irisan lainnya:
// func main() {
// 	myslice1 := []int{1,2,3}
// 	myslice2 := []int{4,5,6}
// 	myslice3 := append(myslice1, myslice2...)

// 	fmt.Printf("myslice3=%v\n", myslice3)
// 	fmt.Printf("length=%d\n", len(myslice3))
// 	fmt.Printf("capacity=%d\n", cap(myslice3))
// }


// -----------------------------------------------------------------------------------------------


// Ubah Panjang Irisan
// Tidak seperti array, dimungkinkan untuk mengubah panjang irisan.
// Contoh ini menunjukkan cara mengubah panjang irisan:
// func main()  {
// 	arr1 := [6] int {9,10,11,12,13,14}		// array
// 	myslice1 := arr1[1:5]	// irisan
// 	fmt.Printf("myslice1 = %v\n", myslice1)
// 	fmt.Printf("length = %d\n", len(myslice1))
// 	fmt.Printf("capacity = %d\n\n", cap(myslice1))

// 	myslice1 = arr1[1:3]	// ubah panjang irisan array
// 	fmt.Printf("myslice1 = %v\n", myslice1)
// 	fmt.Printf("length = %d\n", len(myslice1))
// 	fmt.Printf("capacity = %d\n\n", cap(myslice1))

// 	myslice1 = append(myslice1, 20, 21, 22, 23)	// ubah panjang irisan dengan append()
// 	fmt.Printf("myslice1 = %v\n", myslice1)
// 	fmt.Printf("length = %d\n", len(myslice1))
// 	fmt.Printf("capacity = %d\n\n", cap(myslice1))
// }


// ================================================================================================


// Efisiensi Memori
// Saat menggunakan irisan, Go memuat semua elemen yang mendasarinya ke dalam memori.
// Jika array besar dan Anda hanya membutuhkan beberapa elemen, 
// lebih baik menyalin elemen-elemen tersebut menggunakan fungsi copy().
// Fungsi copy() membuat array baru yang mendasari hanya dengan elemen yang diperlukan untuk irisan. 
// Ini akan mengurangi memori yang digunakan untuk program.
// Fungsi copy() mengambil dua irisan dest dan src, dan menyalin data dari src ke dest. 
// Ini mengembalikan jumlah elemen yang disalin.

// Contoh ini menunjukkan cara menggunakan fungsi copy():
// func main()  {
// 	numbers := [] int {1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18}

// 	// Original slice
// 	fmt.Print("Original slice\n")
// 	fmt.Printf("numbers = %v\n", numbers)
// 	fmt.Printf("panjang = %d\n", len(numbers))
// 	fmt.Printf("kapasitas = %d\n\n", cap(numbers))

// 	// Create copy with only needed numbers
// 	neededNumbers := numbers[:len(numbers)-10]
// 	numbersCopy := make([] int, len(neededNumbers))
// 	copy(numbersCopy, neededNumbers)

// 	fmt.Print("New slice\n")
// 	fmt.Printf("numbers Copy= %v\n", numbersCopy)
// 	fmt.Printf("panjang = %d\n", len(numbersCopy))
// 	fmt.Printf("kapasitas = %d\n\n", cap(numbersCopy))
// }

// Kapasitas irisan baru sekarang lebih kecil 
// dari kapasitas irisan asli karena array baru yang mendasarinya lebih kecil.