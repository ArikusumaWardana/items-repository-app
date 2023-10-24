package main

// import package
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Argumen

// Argumen Tambah
func tambahBarangArgs(args []string) {

	// mengambil argumen nama, harga, stok
	nama := args[1]
	harga := args[2]
	stok := args[3]

	// memasukan data ke dalam slice
	dataBarang := []string{nama, harga, stok}

	// Membuka file atau membuat file baru dan menambahkan data barang
	file, err := os.OpenFile("data-barang.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	// file ditutup setelah fungsinya selesai dieksekusi
	defer file.Close()

	// memisahkan data-data di dalam slice menggunakan |
	_, err = file.WriteString(strings.Join(dataBarang, " | ") + "\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	// menampilkan pesan berhasil jika data berhasil ditambahkan ke dalam file txt
	fmt.Printf("Berhasil menambahkan barang dengan nama %s, harga %s, dan stok %s.", nama, harga, stok)
}

// Fungsi tampilanMenu menampilkan tampilan menu aplikasi
func tampilanMenu() {

	fmt.Println("┌────────────────────────────────────────────────┐")
	fmt.Println("|   Repository Barang PT Maju Kena Mundur Kena   │")
	fmt.Println("├────────────────────────────────────────────────┤")
	fmt.Println("│   Masukkan pilihan Anda:                       │")
	fmt.Println("|                                                |")
	fmt.Println("│   1. Input Barang                              │")
	fmt.Println("│   2. Lihat Daftar Barang                       │")
	fmt.Println("│   3. Cari Barang                               │")
	fmt.Println("│   4. Tentang Aplikasi                          │")
	fmt.Println("│   5. Keluar                                    │")
	fmt.Println("|                                                |")
	fmt.Println("└────────────────────────────────────────────────┘")

}

// Fungsi inputBarang mengambil input dari pengguna dan menyimpannya di file
func inputBarang() {

	// membuat scanner untuk membaca data dari os.Stdin
	readerData := bufio.NewReader(os.Stdin)

	fmt.Println("┌────────────────────────────────────────────────┐")
	fmt.Println("│    Aplikasi Barang PT Maju Kena Mundur Kena    │")
	fmt.Println("└────────────────────────────────────────────────┘")
	fmt.Printf("\n")
	fmt.Println("     Input Barang ")
	fmt.Printf("\n")

	// Mengambil input nama barang
	fmt.Print("    Masukan nama barang  : ")
	nama, _ := readerData.ReadString('\n')
	nama = strings.TrimSpace(nama)

	// Mengambil input harga barang
	fmt.Print("    Masukan harga barang : ")
	harga, _ := readerData.ReadString('\n')
	harga = strings.TrimSpace(harga)

	// Mengambil input stok barang
	fmt.Print("    Masukan stok barang  : ")
	stok, _ := readerData.ReadString('\n')
	stok = strings.TrimSpace(stok)

	fmt.Printf("\n")

	fmt.Println("└────────────────────────────────────────────────┘")

	// memasukan data ke dalam slice
	dataBarang := []string{nama, harga, stok}

	// Membuka file atau membuat file baru dan menambahkan data barang
	file, err := os.OpenFile("data-barang.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		// Menangani kesalahan jika ada saat membuka atau membuat file
		fmt.Println(err)
		return
	}
	// menunda file ditutup setelah fungsinya selesai dieksekusi
	defer file.Close()

	// Menulis dataBarang ke dalam file
	_, err = file.WriteString(strings.Join(dataBarang, " | ") + "\n")
	if err != nil {
		// Menangani kesalahan jika ada saat menulis ke file
		fmt.Println(err)
		return
	}

	// menampilkan pesan sukses jika data berhasil ditulis ke dalam file
	fmt.Println("Data barang berhasil diinputkan!")
	fmt.Printf("\n")

}

// Fungsi showBarang menampilkan daftar barang yang sudah disimpan di file
func showBarang() {

	// Membuka file data-barang.txt untuk dibaca
	fileTxt, err := os.Open("data-barang.txt")
	if err != nil {
		// Menangani kesalahan jika ada saat membuka file
		fmt.Println(err)
		return
	}
	// menunda fileTxt ditutup setelah fungsinya selesai dieksekusi
	defer fileTxt.Close()

	// membuat scanner untuk membaca data dari file
	scanFile := bufio.NewScanner(fileTxt)
	fmt.Println("Lihat Daftar Barang")
	fmt.Println("┌────────────────────────────────────────────────┐")
	fmt.Println("|               LIHAT DAFTAR BARANG              |")
	fmt.Println("└────────────────────────────────────────────────┘")
	fmt.Printf("\n")

	// Menampilkan data barang dengan nomor urut
	for i := 1; scanFile.Scan(); i++ {
		data := scanFile.Text()
		fmt.Print("    ")
		fmt.Printf("%d. %s\n", i, data)
	}
	fmt.Println("└────────────────────────────────────────────────┘")
	fmt.Printf("\n")
}

// Fungsi search mencari barang berdasarkan nama
func search() {
	// membuat scanner untuk membaca data dari os.Stdin
	readerData := bufio.NewReader(os.Stdin)
	fmt.Println("┌────────────────────────────────────────────────┐")
	fmt.Println("|                   CARI BARANG                  |")
	fmt.Println("└────────────────────────────────────────────────┘")

	fmt.Print("  > Masukan nama barang: ")
	nama, _ := readerData.ReadString('\n')
	nama = strings.TrimSpace(nama)

	// Memanggil fungsi searchBarang dengan parameter nama
	searchBarang(nama)
}

// Fungsi searchBarang melakukan pencarian dan menampilkan hasilnya
func searchBarang(namaBarang string) {
	// Membuka file-barang.txt untuk dibaca
	fileTxt, err := os.Open("data-barang.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fileTxt ditutup setelah fungsinya selesai dieksekusi
	defer fileTxt.Close()

	// membuat scanner untuk membaca data dari file
	scanBarang := bufio.NewScanner(fileTxt)
	// menandai apakah sudah ditemukan setidaknya satu baris data yang sesuai dengan kata kunci pencarian
	found := false
	i := 1
	// Melakukan iterasi melalui setiap baris file
	for scanBarang.Scan() {
		// Membaca data dari file
		data := scanBarang.Text()
		// Memeriksa apakah data mengandung kata yang dicari
		if matchIgnoreCase(data, namaBarang) {
			// Menampilkan data jika ditemukan
			if !found {
				fmt.Println("    Data ditemukan:")
				found = true
			}
			fmt.Print("    ")
			fmt.Printf("%d. %s\n", i, data)
			i++
		}
	}
	// Menangani kesalahan jika ada saat membaca file
	if err := scanBarang.Err(); err != nil {
		fmt.Println(err)
	}

	// Menampilkan pesan jika data tidak ditemukan
	if !found {
		fmt.Println("Data tidak ditemukan!")
	}
	fmt.Println("└────────────────────────────────────────────────┘")
}

// Fungsi matchIgnoreCase mencocokkan string tanpa memperhatikan huruf besar/kecil
func matchIgnoreCase(haystack, needle string) bool {
	for _, kata := range strings.Fields(needle) {
		if !strings.Contains(strings.ToLower(haystack), strings.ToLower(kata)) {
			return false
		}
	}
	return true
}

// Fungsi detailApp menampilkan informasi tentang aplikasi
func detailApp() {

	fmt.Println("┌────────────────────────────────────────────────┐")
	fmt.Println("|   Repository Barang PT Maju Kena Mundur Kena   │")
	fmt.Println("├────────────────────────────────────────────────┤")
	fmt.Println("│   TENTANG APLIKASI                             │")
	fmt.Println("|                                                |")
	fmt.Println("│   NAMA PEMBUAT: KADEK AGUS ARIKUSUMA WARDANA   │")
	fmt.Println("│   NIM         : 2301020033                     │")
	fmt.Println("│   KELAS       : PAGI2                          │")
	fmt.Println("|                                                |")
	fmt.Println("└────────────────────────────────────────────────┘")

}

// Fungsi pilihMenu menangani pilihan menu dari pengguna
func pilihMenu() {

	var menu string

	// Mengambil nomor input dari user
	fmt.Print("    > Masukan nomor menu: ")
	fmt.Scanln(&menu)
	fmt.Printf("\n")

	switch menu {
	// Jika input user adalah "1", panggil fungsi inputBarang()
	case "1":
		inputBarang()
	// Jika input user adalah "2", panggil fungsi showBarang()
	case "2":
		showBarang()
	// Jika input user adalah "3", panggil fungsi search()
	case "3":
		search()
	// Jika input user adalah "4", panggil fungsi detailApp()
	case "4":
		detailApp()
	// Jika input user adalah "5", maka aplikasi akan dihentikan
	case "5":
		fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
		os.Exit(0)
	// Jika input user selain dari 1 - 5, maka akan tampil pesan alert
	default:
		fmt.Println("Nomor menu yang dimaksukan salah! silahkan coba lagi")
	}

}

// fungsi utama
func main() {

	// mengecek perkondisian apakah panjang argumen yg diinputkan lebih dari 1 argumen
	if len(os.Args) > 1 {

		arg := os.Args[1]

		switch arg {
		// Jika ada argumen "tambah", panggil fungsi tambahBarangArgs(), dengan argumen yang diberikan
		case "tambah":
			if len(os.Args) >= 4 {
				tambahBarangArgs(os.Args[1:])
			} else {
				fmt.Println("Penggunaan : go run main.go tambah (nama barang) (harga) (stok)")
			}
		// Jika ada argumen "show", panggil fungsi showBarang()
		case "show":
			showBarang()

		// Jika ada argumen "search", panggil fungsi searchBarang(), dengan argumen yang diberikan
		case "search":
			if len(os.Args) > 2 {
				searchBarang(os.Args[2])
			} else {
				fmt.Println("Penggunaan : go run main.go search (nama barang)")
			}

		// Jika ada argumen "about", panggil fungsi detailApp()
		case "detail":
			detailApp()
		// Jika argumen selain dari yang diatas, maka akan tampil pesan alert
		default:
			fmt.Println("Penggunaan : go run main.go [tambah/show/search/detail]")
		}

	} else {
		// Jika tidak ada argumen, jalankan aplikasi seperti biasa
		for {
			tampilanMenu() // memanggil fungsi tampilanMenu
			pilihMenu()    // memanggil fungsi piihMenu

			// menginstruksikan user untuk menekan enter terlebih dahulu sebelum menampilkan menu
			fmt.Print("Tekan Enter untuk melanjutkan...")
			fmt.Scanln()
		}
	}

}
