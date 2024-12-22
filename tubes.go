package main

import (
	"fmt"
	"strings"
	"time"
)

const NMAX int = 5000

type User struct {
	Username string
	Password string
	Role     string
}

type tabUser struct {
	Users [NMAX]User
	n     int
}

type Obat struct {
	Nama, Kode, Pabrikan, TanggalKadaluwarsa string
	Kategori                                 string
	JumlahStok, Harga                        int
}

type tabObat struct {
	DaftarObat [NMAX]Obat
	n          int
}

func register(U *tabUser) {
	var username, password string

	fmt.Print("Masukkan username: ")
	fmt.Scanln(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scanln(&password)

	U.Users[U.n] = User{Username: username, Password: password, Role: "buyer"}
	U.n++
	fmt.Println("Registrasi berhasil.")
}

func initAdmin(U *tabUser) {
	admin := User{Username: "admin", Password: "bukanadmin", Role: "admin"}
	U.Users[U.n] = admin
	U.n++
}

func login(U *tabUser) *User {
	var username, password string

	fmt.Print("Masukkan username: ")
	fmt.Scanln(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scanln(&password)

	for i := 0; i < U.n; i++ {
		if U.Users[i].Username == username && U.Users[i].Password == password {
			fmt.Println("Login berhasil.")
			return &U.Users[i]
		}
	}
	fmt.Println("Username atau password salah.")
	return nil
}

func adminMenu(database *tabObat) {
	for {
		fmt.Println("\nPilih operasi yang ingin dilakukan:")
		fmt.Println("1. Tambahkan Obat")
		fmt.Println("2. Cari Obat")
		fmt.Println("3. Edit Obat")
		fmt.Println("4. Hapus Obat")
		fmt.Println("5. Tampilkan Obat Berdasarkan Harga")
		fmt.Println("6. Tampilkan Obat Berdasarkan Tanggal Kadaluwarsa")
		fmt.Println("7. Keluar")
		fmt.Print("Pilihan: ")
		var pilihan string
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			tambahObat(database)
		case "2":
			cariObat(*database)
		case "3":
			editObat(database)
		case "4":
			hapusObat(database)
		case "5":
			tampilkanObatharga(database)
		case "6":
			fmt.Println("\nPilih urutan yang diinginkan:")
			fmt.Println("1. Dari paling awal hingga paling akhir")
			fmt.Println("2. Dari paling akhir hingga paling awal")
			fmt.Print("Pilihan: ")
			var opsi string
			fmt.Scanln(&opsi)
			if opsi == "1" {
				tampilkanObatKadaluarsa(database)
			} else if opsi == "2" {
				tampilkanObatKadaluarsaDescending(database)
			} else {
				fmt.Println("Pilihan tidak valid.")
			}
		case "7":
			fmt.Println("Keluar dari menu admin.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
		}
	}
}

func buyerMenu(database *tabObat) {
	for {
		fmt.Println("\nPilih operasi yang ingin dilakukan:")
		fmt.Println("1. Cari Obat")
		fmt.Println("2. Pesan Obat")
		fmt.Println("3. Keluar")
		fmt.Print("Pilihan: ")
		var pilihan string
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			cariObat(*database)
		case "2":
			pesanObat(database)
		case "3":
			fmt.Println("Keluar dari menu pembeli.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
		}
	}
}

func tambahObat(A *tabObat) {
	for A.n < NMAX {
		var obat Obat
		fmt.Printf("Masukkan obat ke-%d\n", A.n+1)
		fmt.Print("Nama: ")
		fmt.Scanln(&obat.Nama)
		if strings.ToLower(obat.Nama) == "selesai" {
			fmt.Println("Data Berhasil Ditambahkan!")
			return
		}
		fmt.Print("Kode: ")
		fmt.Scanln(&obat.Kode)
		fmt.Print("Harga: ")
		fmt.Scanln(&obat.Harga)
		fmt.Print("Jumlah Stok: ")
		fmt.Scanln(&obat.JumlahStok)
		fmt.Print("Tanggal Kadaluwarsa (format: YYYY-MM-DD): ")
		fmt.Scanln(&obat.TanggalKadaluwarsa)
		fmt.Print("Kategori: ")
		fmt.Scanln(&obat.Kategori)
		fmt.Print("Pabrikan: ")
		fmt.Scanln(&obat.Pabrikan)

		A.DaftarObat[A.n] = obat
		A.n++
	}
}

func cariObat(A tabObat) {
	var input string
	fmt.Print("Masukkan kode atau nama obat: ")
	fmt.Scanln(&input)
	ketemu := false

	for i := 0; i < A.n; i++ {
		if strings.ToLower(A.DaftarObat[i].Kode) == strings.ToLower(input) || strings.ToLower(A.DaftarObat[i].Nama) == strings.ToLower(input) {
			ketemu = true
			fmt.Println("Data Obat:")
			fmt.Println("Nama:", A.DaftarObat[i].Nama)
			fmt.Println("Kode:", A.DaftarObat[i].Kode)
			fmt.Println("Harga:", A.DaftarObat[i].Harga)
			fmt.Println("Jumlah Stok:", A.DaftarObat[i].JumlahStok)
			fmt.Println("Tanggal Kadaluwarsa:", A.DaftarObat[i].TanggalKadaluwarsa)
			fmt.Println("Kategori:", A.DaftarObat[i].Kategori)
			fmt.Println("Pabrikan:", A.DaftarObat[i].Pabrikan)
			return
		}
	}

	if !ketemu {
		fmt.Println("Obat dengan kode atau nama tersebut tidak ditemukan.")
		return
	}
}

func editObat(A *tabObat) {
	var input string
	fmt.Print("Masukkan nama atau kode obat yang ingin diedit: ")
	fmt.Scanln(&input)

	found := false
	for i := 0; i < A.n; i++ {
		if strings.Contains(strings.ToLower(A.DaftarObat[i].Nama), strings.ToLower(input)) || strings.ToLower(A.DaftarObat[i].Kode) == strings.ToLower(input) {
			fmt.Println("Data Obat yang akan diedit:")
			fmt.Println("Nama:", A.DaftarObat[i].Nama)
			fmt.Println("Kode:", A.DaftarObat[i].Kode)
			fmt.Println("Harga:", A.DaftarObat[i].Harga)
			fmt.Println("Jumlah Stok:", A.DaftarObat[i].JumlahStok)
			fmt.Println("Tanggal Kadaluwarsa:", A.DaftarObat[i].TanggalKadaluwarsa)
			fmt.Println("Kategori:", A.DaftarObat[i].Kategori)
			fmt.Println("Pabrikan:", A.DaftarObat[i].Pabrikan)

			found = true

			var field string
			fmt.Println("\nPilih data yang ingin diubah:")
			fmt.Println("1. Nama")
			fmt.Println("2. Harga")
			fmt.Println("3. Jumlah Stok")
			fmt.Println("4. Tanggal Kadaluwarsa")
			fmt.Println("5. Kategori")
			fmt.Println("6. Pabrikan")
			fmt.Print("Pilihan: ")
			fmt.Scanln(&field)

			switch field {
			case "1":
				fmt.Print("Masukkan nama baru: ")
				fmt.Scanln(&A.DaftarObat[i].Nama)
			case "2":
				fmt.Print("Masukkan harga baru: ")
				fmt.Scanln(&A.DaftarObat[i].Harga)
			case "3":
				fmt.Print("Masukkan jumlah stok baru: ")
				fmt.Scanln(&A.DaftarObat[i].JumlahStok)
			case "4":
				fmt.Print("Masukkan tanggal kadaluwarsa baru (format: YYYY-MM-DD): ")
				fmt.Scanln(&A.DaftarObat[i].TanggalKadaluwarsa)
			case "5":
				fmt.Print("Masukkan kategori baru: ")
				fmt.Scanln(&A.DaftarObat[i].Kategori)
			case "6":
				fmt.Print("Masukkan pabrikan baru: ")
				fmt.Scanln(&A.DaftarObat[i].Pabrikan)
			default:
				fmt.Println("Pilihan tidak valid.")
				return
			}

			fmt.Println("Data obat berhasil diupdate.")
			return
		}
	}

	if !found {
		fmt.Println("Obat dengan nama atau kode tersebut tidak ditemukan.")
	}
}

func hapusObat(A *tabObat) {
	sortingAscendingByHarga(A)
	var input string
	fmt.Print("Masukkan nama atau kode obat yang ingin dihapus: ")
	fmt.Scanln(&input)

	index := cariBinary(A.DaftarObat[:A.n], strings.ToLower(input))
	if index != -1 {
		fmt.Println("Data Obat yang akan dihapus:")
		fmt.Println("Nama:", A.DaftarObat[index].Nama)
		fmt.Println("Kode:", A.DaftarObat[index].Kode)
		fmt.Println("Harga:", A.DaftarObat[index].Harga)
		fmt.Println("Jumlah Stok:", A.DaftarObat[index].JumlahStok)
		fmt.Println("Tanggal Kadaluwarsa:", A.DaftarObat[index].TanggalKadaluwarsa)
		fmt.Println("Kategori:", A.DaftarObat[index].Kategori)
		fmt.Println("Pabrikan:", A.DaftarObat[index].Pabrikan)

		var confirm string
		fmt.Print("Apakah Anda yakin ingin menghapus data obat ini? (ya/tidak): ")
		fmt.Scanln(&confirm)
		if strings.ToLower(confirm) == "ya" {
			for i := index; i < A.n-1; i++ {
				A.DaftarObat[i] = A.DaftarObat[i+1]
			}
			A.n--
			fmt.Println("Data obat berhasil dihapus.")
		} else {
			fmt.Println("Penghapusan dibatalkan.")
		}
	} else {
		fmt.Println("Obat dengan nama atau kode tersebut tidak ditemukan.")
	}
}

func pesanObat(A *tabObat) {
	var input string
	fmt.Print("Masukkan nama atau kode obat yang ingin dipesan: ")
	fmt.Scanln(&input)

	found := false
	for i := 0; i < A.n; i++ {
		if strings.Contains(strings.ToLower(A.DaftarObat[i].Nama), strings.ToLower(input)) || strings.ToLower(A.DaftarObat[i].Kode) == strings.ToLower(input) {
			found = true
			fmt.Println("Data Obat:")
			fmt.Println("Nama:", A.DaftarObat[i].Nama)
			fmt.Println("Kode:", A.DaftarObat[i].Kode)
			fmt.Println("Harga:", A.DaftarObat[i].Harga)
			fmt.Println("Jumlah Stok:", A.DaftarObat[i].JumlahStok)
			fmt.Println("Tanggal Kadaluwarsa:", A.DaftarObat[i].TanggalKadaluwarsa)
			fmt.Println("Kategori:", A.DaftarObat[i].Kategori)
			fmt.Println("Pabrikan:", A.DaftarObat[i].Pabrikan)

			var jumlahPesan int
			fmt.Print("Masukkan jumlah yang ingin dipesan: ")
			fmt.Scanln(&jumlahPesan)

			if jumlahPesan > 0 && jumlahPesan <= A.DaftarObat[i].JumlahStok {
				A.DaftarObat[i].JumlahStok -= jumlahPesan
				fmt.Println("Obat berhasil dipesan.")
				fmt.Printf("Total harga: %d\n", jumlahPesan*A.DaftarObat[i].Harga)
			} else {
				fmt.Println("Jumlah pesan tidak valid atau stok tidak mencukupi.")
			}
			return
		}
	}

	if !found {
		fmt.Println("Obat dengan nama atau kode tersebut tidak ditemukan.")
	}
}
func tampilkanObatMahal(A tabObat) {
	sortingAscendingByHarga(&A)
	fmt.Println("\nDaftar Obat Berdasarkan Harga tertinggi ke terendah:")
	for i := 0; i < A.n; i++ {
		fmt.Printf("%d. Nama: %s, Harga: %d\n", i+1, A.DaftarObat[i].Nama, A.DaftarObat[i].Harga)
	}
}

func tampilkanObatHargaMurah(A tabObat) {
	sortingAscendingByHarga(&A)
	fmt.Println("Daftar Obat Berdasarkan Harga termurah ke termahal:")
	for i := 0; i < A.n; i++ {
		fmt.Printf("%d. Nama: %s, Harga: %d\n", i+1, A.DaftarObat[i].Nama, A.DaftarObat[i].Harga)
	}
}

func tampilkanObatharga(A *tabObat) {
	fmt.Println("\nPilih opsi pengurutan harga:")
	fmt.Println("1. Tertinggi ke Terendah")
	fmt.Println("2. Termurah ke Termahal")
	fmt.Print("Pilihan: ")
	var pilihan string
	fmt.Scanln(&pilihan)

	switch pilihan {
	case "1":
		tampilkanObatMahal(*A)
	case "2":
		tampilkanObatHargaMurah(*A)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func tampilkanObatKadaluarsa(A *tabObat) {
	insertionSortByTanggalKadaluarsa(A.DaftarObat[:A.n])
	fmt.Println("Data obat berdasarkan tanggal kadaluwarsa (terlama ke terbaru):")
	for i := 0; i < A.n; i++ {
		fmt.Println("Nama:", A.DaftarObat[i].Nama)
		fmt.Println("Tanggal Kadaluwarsa:", A.DaftarObat[i].TanggalKadaluwarsa)
	}
}

func tampilkanObatKadaluarsaDescending(A *tabObat) {
	insertionSortByTanggalKadaluarsaDescending(A.DaftarObat[:A.n])
	fmt.Println("Data obat berdasarkan tanggal kadaluwarsa (terbaru ke terlama):")
	for i := 0; i < A.n; i++ {
		fmt.Println("Nama:", A.DaftarObat[i].Nama)
		fmt.Println("Tanggal Kadaluwarsa:", A.DaftarObat[i].TanggalKadaluwarsa)
	}
}

func sortingAscendingByNama(A *tabObat) {
	for i := 1; i < A.n; i++ {
		key := A.DaftarObat[i]
		j := i - 1
		for j >= 0 && strings.ToLower(A.DaftarObat[j].Nama) < strings.ToLower(key.Nama) {
			A.DaftarObat[j+1] = A.DaftarObat[j]
			j = j - 1
		}
		A.DaftarObat[j+1] = key
	}
}

func sortingAscendingByHarga(A *tabObat) {
	for i := 1; i < A.n; i++ {
		key := A.DaftarObat[i]
		j := i - 1
		for j >= 0 && A.DaftarObat[j].Harga < key.Harga {
			A.DaftarObat[j+1] = A.DaftarObat[j]
			j = j - 1
		}
		A.DaftarObat[j+1] = key
	}
}

func sortingDescendingByHarga(A *tabObat) {
	for i := 1; i < A.n; i++ {
		key := A.DaftarObat[i]
		j := i - 1
		for j >= 0 && A.DaftarObat[j].Harga > key.Harga {
			A.DaftarObat[j+1] = A.DaftarObat[j]
			j = j - 1
		}
		A.DaftarObat[j+1] = key
	}
}

func insertionSortByTanggalKadaluarsa(obatList []Obat) {
	n := len(obatList)
	for i := 1; i < n; i++ {
		key := obatList[i]
		keyDate, _ := time.Parse("2006-01-02", key.TanggalKadaluwarsa)
		j := i - 1
		for j >= 0 {
			jDate, _ := time.Parse("2006-01-02", obatList[j].TanggalKadaluwarsa)
			if jDate.After(keyDate) {
				obatList[j+1] = obatList[j]
				j = j - 1
			} else {
				return
			}
		}
		obatList[j+1] = key
	}
}

func insertionSortByTanggalKadaluarsaDescending(obatList []Obat) {
	n := len(obatList)
	for i := 1; i < n; i++ {
		key := obatList[i]
		keyDate, _ := time.Parse("2006-01-02", key.TanggalKadaluwarsa)
		j := i - 1
		for j >= 0 {
			jDate, _ := time.Parse("2006-01-02", obatList[j].TanggalKadaluwarsa)
			if jDate.Before(keyDate) {
				obatList[j+1] = obatList[j]
				j = j - 1
			} else {
				return
			}
		}
		obatList[j+1] = key
	}
}

func cariBinary(obatList []Obat, input string) int {
	low := 0
	high := len(obatList) - 1

	for low <= high {
		mid := (low + high) / 2
		if strings.ToLower(obatList[mid].Nama) == input || strings.ToLower(obatList[mid].Kode) == input {
			return mid
		} else if strings.ToLower(obatList[mid].Nama) > input {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	var database tabObat
	var users tabUser

	initAdmin(&users)

	for {
		var choice string
		fmt.Println("Selamat datang di sistem manajemen apotek.")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			user := login(&users)
			if user != nil {
				if user.Role == "admin" {
					adminMenu(&database)
				} else {
					buyerMenu(&database)
				}
			}
		case "2":
			register(&users)
		case "3":
			fmt.Println("Terima kasih telah menggunakan sistem manajemen apotek.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
		}
	}
}
