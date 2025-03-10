package main
import "fmt"

// Array digunakan untuk menyimpan beberapa nilai dari jenis yang sama dalam satu variabel, 
// alih-alih mendeklarasikan variabel terpisah untuk setiap nilai.
// Panjangnya menentukan jumlah elemen untuk disimpan dalam array. 
// Di Go, array memiliki panjang tetap. 
// Panjang array didefinisikan oleh angka atau disimpulkan 
// (berarti bahwa kompiler memutuskan panjang array, berdasarkan jumlah nilai).

// Contoh ini mendeklarasikan dua array (arr1 dan arr2) dengan panjang yang ditentukan:
// func main()  {
// 	var arr1 = [3] int {1,2,3}
// 	arr2 := [5] int {4,5,6,7,8}

// 	fmt.Println(arr1, "\n", arr2)
// }

// Contoh ini mendeklarasikan dua array (arr1 dan arr2) dengan panjang yang disimpulkan:
// func main()  {
// 	arr1 := [...] int {1,2,3}
// 	arr2 := [...] int {4,5,6,7,8}

// 	fmt.Println(arr1, "\n", arr2)
// }

// Contoh ini mendeklarasikan array string:
// func main()  {
// 	cars := [4] string {"BMW", "Ford", "Toyota", "Mazda"}
// 	fmt.Print(cars)
// }


// ==================================================================================================


// Mengakses elemen array
// Anda dapat mengakses elemen array tertentu dengan mengacu pada nomor indeks.
// Di Go, indeks array dimulai dari 0. 
// Itu berarti bahwa [0] adalah elemen pertama, [1] adalah elemen kedua, dll.

// Contoh ini menunjukkan cara mengakses elemen pertama dan ketiga dalam array prices:
// func main()  {
// 	prices := [...] int {10,5,23,20,90,92}
// 	fmt.Println(prices[0])
// 	fmt.Println(prices[2])
// 	fmt.Println(prices[5])
// 	fmt.Println(prices[1])
// }


// ====================================================================================================


// Ubah elemen array
// Anda juga dapat mengubah nilai elemen array tertentu dengan mengacu pada nomor indeks.

// Contoh ini menunjukkan cara mengubah nilai elemen ketiga dalam array prices:
// func main() {
// 	prices := [3]int{10,20,30}

// 	prices[2] = 50
// 	fmt.Println(prices)
// }


// ====================================================================================================


// Inisialisasi Array
// Jika array atau salah satu elemennya belum diinisialisasi dalam kode, itu diberi nilai default dari jenisnya.
// Nilai default untuk int adalah 0, dan nilai default untuk string adalah "".

// contoh
// func main() {
// 	arr1 := [5]int{} 			//not initialized
// 	arr2 := [5]int{1,2} 		//partially initialized
// 	arr3 := [5]int{1,2,3,4,5} 	//fully initialized

// 	fmt.Println(arr1)
// 	fmt.Println(arr2)
// 	fmt.Println(arr3)
// }


// ===============================================================================================


// Inisialisasi Hanya Elemen Tertentu
// Dimungkinkan untuk menginisialisasi hanya elemen tertentu dalam array.

// Contoh ini hanya menginisialisasi elemen kedua dan ketiga dari array:
// func main() {
// 	arr1 := [5]int{1:10,2:40}

// 	fmt.Println(arr1)

// 	// Array di atas memiliki 5 elemen.
// 	// 1:10 berarti: tetapkan 10 ke indeks array 1 (elemen kedua).
// 	// 2:40 berarti: tetapkan 40 ke indeks array 2 (elemen ketiga).
// }


// ===========================================================================================


// Temukan Panjang Array
// Fungsi len() digunakan untuk mencari panjang array:
func main() {
	arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	arr2 := [...]int{1,2,3,4,5,6}

	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
}