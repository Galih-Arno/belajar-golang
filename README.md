# Belajar Golang Dasar

## Apa itu Go?

* Go adalah bahasa pemrograman yang cepat, dinamis, statis, dan dikompilasi serta open source.
* Go dikembangkan di Google oleh Robert Griesemer, Rob Pike, dan Ken Thompson pada 2007.
* Go dapat digunakan untuk membuat aplikasi berkinerja tinggi dengan sintaks yang mirip dengan C++.
* Go mendukung concurrency, memory management, fast run time, dan compilation time.

## Go digunakan untuk:

* Pengembangan web (sisi server)
* Program berbasis jaringan
* Pengembangan aplikasi cross-platform dan cloud-native

> concurrency melakukan beberapa hal secara tidak berurutan, atau pada saat yang sama, tanpa mempengaruhi hasil akhir.

Teks editor yang biasanya digunakan:
* [Visual Studi Code](https://code.visualstudio.com/)
* [GoLand](https://www.jetbrains.com/go/promo/?source=google&medium=cpc&campaign=APAC_en_ASIA_GoLand_Search&term=go%20editor&content=545953842864&gad_source=1&gclid=Cj0KCQjwm7q-BhDRARIsACD6-fWx4c7PjTY4PQcK_sR58WX0AXptIF_Qt994dwfOnqWqLpoD-jjRrBkaArOnEALw_wcB)

## Install GO

kamu bisa menginstallnya di [GO](https://go.dev/dl/), kemudian ikuti petunjuk yang terkait dengan sistem operasimu. Untuk memeriksa apakah Go berhasil diinstall, kamu bisa menjalankan perintah berikut di terminal:
```shell
    go version
```

Mantap, kamu pengen yang **lebih mudah dan singkat** ya? Oke, aku coba sederhanakan banget, kayak ngobrol aja. Ini dia:

---

## 🚀 Cara Super Mudah Bikin Program Pertama di Golang

### 1. **Pasang Dulu VS Code + Extension Go**
- **Buka VS Code.**
- Cari logo kotak-kotak (bagian kiri) → Klik.
- Ketik **"Go"** di kolom cari → Pilih yang **Go (by Google)** → Klik **Install**.
- Setelah itu, tekan **Ctrl + Shift + P** → ketik **"Go: Install/Update Tools"** → pilih semua → **OK**.

🎉 **VS Code siap untuk coding Go!**

---

### 2. **Bikin Proyek Baru**
- Buka **terminal** di VS Code (klik `Terminal > New Terminal`).
- Ketik ini (buat ngasih tahu Go, "eh aku mau bikin proyek baru nih"):
  ```bash
  go mod init example.com/hello
  ```
💡 **Anggap aja ini kayak "setup project", step wajib aja dulu.**

---

### 3. **Tulis Kode Go Pertama**
- Klik `File > New File`.
- Tulis ini:
  ```go
  package main
  import "fmt"

  func main() {
    fmt.Println("Hello World!")
  }
  ```
- Simpan file ini, kasih nama: `helloworld.go`.

---

### 4. **Jalankan Kode Go**
- Masih di terminal, ketik:
  ```bash
  go run helloworld.go
  ```
- ✅ Nanti muncul tulisan:
  ```
  Hello World!
  ```

---

### 🎉 Udah Jadi! 🎉
Kamu **berhasil bikin program Go pertama**.  

---

### ✨ Intinya:
| Langkah                   | Apa yang Dilakukan                                    |
|-------------------------|-----------------------------------------------------|
| Install VS Code + Go     | Biar bisa ngoding Go di VS Code                     |
| `go mod init`            | Setup project                                       |
| Tulis file `helloworld.go` | Isi kode program "Hello World"                     |
| `go run helloworld.go`  | Jalankan program                                    |
