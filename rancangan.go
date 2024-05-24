package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
)

const NMAX int = 256

type tanggal struct {
	tanggal, bulan, tahun int
}

type pasien struct {
	id     int
	nama   string
	waktu  tanggal
	paket  string
	ada    bool
	biaya  int
	umur   int
}

type paketMCU struct {
	id    int
	nama  string
	biaya int
	ada   bool
}

type registrasi struct {
	nama, UserID, password, job string
}

type tabPasien [NMAX]pasien
type tabPaket [NMAX]paketMCU
type tabRegistrasi [NMAX]registrasi

func main() {
	var a tabRegistrasi
	var b tabPasien
	var c tabPaket
	var n, m, p int
	fmt.Println("======================================================")
	fmt.Println("|              \x1b[32mSELAMAT DATANG\x1b[0m               |")
	fmt.Println("|             \x1b[32mAVICENNA HOSPITAL\x1b[0m             |")
	fmt.Println("======================================================")
	jedawaktu(3)
	cls()
	loading()
	menuawal(&a,&n,&b,&c,&m,&p)
}

func menuawal(pengguna *tabRegistrasi, nPengguna *int, a *tabPasien, b *tabPaket, m, p *int) {
	var i int
	headerregistrasi21()
	fmt.Scan(&i)
	switch i {
	case 1:
		buatregistrasi(&*pengguna,&*nPengguna)
		menuawal(&*pengguna,&*nPengguna,&*a,&*b,&*m,&*p)
	case 2:
		login(*pengguna,*nPengguna,&*a,&*b,&*m,&*p)
	case 3:
		logout()
	}
}

func loading() {
	cls()
	fmt.Print("Loading")
	jedawaktu(1)
	fmt.Print(".")
	jedawaktu(1)
	fmt.Print(".")
	jedawaktu(1)
	fmt.Print(".")
	cls()
}

func menuutama(a *tabPasien, b *tabPaket, m, p *int) {
	var pilihan int

	Header()
	fmt.Println("Pilihan Anda: ")
	fmt.Scan(&pilihan)
	for pilihan != 6 {
		switch pilihan {
		case 1:
			menuInputData(a, b, m, p)
		case 2:
			menuUbahData(a, b, m, p)
		case 3:
			deleteData(a, b, m, p)
		case 4:
			menampilkanData(a, b, m, p)
		case 5:
			sortingData(a, b, m, p)
		default:
			fmt.Println("======Pilihan Tidak Sesuai======")
		}
		Header()
		fmt.Println("Pilihan Anda:")
		fmt.Scan(&pilihan)
	}
	exit()
}

func buatregistrasi(pengguna *tabRegistrasi, nPengguna *int) {
	var valid, stop bool
	var pilih, target string
	var i int
	valid = false
	if *nPengguna < NMAX {
		fmt.Println("==========DATA REGISTRASI==========")
		fmt.Print("Masukkan Nama Anda: ")
		fmt.Scan(&pengguna[*nPengguna].nama)
		fmt.Print("Pilih Status Anda: ")
		fmt.Print("\n1. Petugas Rumah Sakit\n2. Laboratorium\n(1/2)? ")
		for !valid {
			fmt.Scan(&pilih)
			if pilih == "1" {
				pengguna[*nPengguna].job = "Petugas Rumah Sakit"
				valid = true
			} else if pilih == "2" {
				pengguna[*nPengguna].job = "Laboratorium"
				valid = true
			} else {
				fmt.Println("\x1b[31mPilihan Anda Tidak Valid\x1b[0m")
			}
		}
		fmt.Print("Masukkan UserID Anda: ")
		for !stop {
			valid = false
			fmt.Scan(&pengguna[*nPengguna].UserID)
			target = pengguna[*nPengguna].UserID
			for i = 0; i < *nPengguna; i++ {
				if target == pengguna[i].UserID {
					valid = true
				}
			}
			if valid {
				fmt.Println("UserID Tidak VALID")
				fmt.Print("Masukkan UserID: ")
			} else {
				fmt.Print("Buat Password Anda: ")
				fmt.Scan(&pengguna[*nPengguna].password)
			}
			if valid == false {
				stop = true
			}
		}
		*nPengguna++
	}
	headerregistrasi()
}

func login(pengguna tabRegistrasi, nPengguna int, a *tabPasien, b *tabPaket, m, p *int) { //?Procedure untuk menampilkan login
	var username, password, target string
	var stop bool
	var i int

	stop = false
	fmt.Print("Masukkan UserID Anda: ")
	fmt.Scan(&username)
	target = username
	for i = 0; i < NMAX; i++ {
		if pengguna[i].UserID == target {
			fmt.Print("Masukkan Password: ")
			fmt.Scan(&password)
			target = password
			jedawaktu(3)
			if target == pengguna[i].password {
				stop = true
			} else {
				for target != pengguna[i].password {
					fmt.Println("=====Password Invalid=====")
					fmt.Print("Masukkan Password: ")
					fmt.Scan(&target)
					jedawaktu(3)
				}
				stop = true
			}
		}
	}
	if stop == false {
		fmt.Println()
		cls()
		headerregistrasi2()
		fmt.Scan(&i)
		if i == 1 {
			buatregistrasi(&pengguna, &nPengguna)
		} else if i == 2 {
			login(pengguna, nPengguna, a, b, m, p)
		} else if i == 3 {
			exit()
		}
	}
	cls()
	headerlogin()
	fmt.Scan(&i)
	if i == 1 {
		loading()
		menuutama(a, b, m, p)
	} else if i == 2 {
		logout()
	} else if i == 3 {
		exit()
	}
}

func logout() {
	cls()
	headerregistrasi()
}

func exit() {
	cls()
	fmt.Print("\x1b[32m=======KEPUASAN ANDA ADALAH PRIORITAS KAMI=======\x1b[0m")
}

func menuInputData(a *tabPasien, b *tabPaket, m, p *int) {
	var pilihan int
	var i int
	var stop bool

	stop = false
	HeaderinputData()
	fmt.Println("Pilihan Anda: ")
	fmt.Scan(&pilihan)
	for !stop {
		if pilihan == 1 {
			Headerinput1()
			for i = 0; i < NMAX && !stop; i++ {
				if !a[i].ada {
					fmt.Print("Enter Patient ID: ")
					fmt.Scan(&a[i].id)
					fmt.Print("Enter Patient Name: ")
					fmt.Scan(&a[i].nama)
					fmt.Print("Enter Patient Age: ")
					fmt.Scan(&a[i].umur)
					fmt.Print("Paket registrasi: ")
					fmt.Scan(&a[i].paket)
					paketValid := false
					for j := 0; j < *p; j++ {
						if b[j].nama == a[i].paket && b[j].ada {
							a[i].biaya = b[j].biaya
							paketValid = true
							stop = true
						}
					}
					if !paketValid {
						fmt.Println("\n=== Paket Tidak Tersedia ===")
						stop = false
					} else {
						fmt.Print("Waktu registrasi (yyyy mm dd): ")
						fmt.Scan(&a[i].waktu.tahun, &a[i].waktu.bulan, &a[i].waktu.tanggal)
						a[i].ada = true
						stop = true
					}
				}
			}
		}
		if pilihan == 2 {
			Headerinput12()
			for i = 0; i < NMAX && !stop; i++ {
				if !b[i].ada {
					fmt.Print("Enter package ID: ")
					fmt.Scan(&b[i].id)
					fmt.Print("Enter Package Name: ")
					fmt.Scan(&b[i].nama)
					fmt.Print("Biaya: ")
					fmt.Scan(&b[i].biaya)
					b[i].ada = true
					stop = true
					*p++
				}
			}
		}
		if pilihan == 3 {
			cls()
			HeaderinputData()
		}
		if pilihan == 4 {
			cls()
			exit()
		}
	}
}

func menuUbahData(a *tabPasien, b *tabPaket, m, p *int) {
	var pilihan int

	fmt.Println("Pilihan Anda untuk mengubah data: ")
	fmt.Println("1. Ubah Data Pasien")
	fmt.Println("2. Ubah Data Paket")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		ubahDataPasien(a, *m)
	} else if pilihan == 2 {
		ubahDataPaket(b, *p)
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func ubahDataPasien(a *tabPasien, m int) { //?Procedure untuk mengubah data pasien
	var id, umur, pilihan int
	var nama, paket string
	var tgl, bln, thn int

	fmt.Print("Masukkan ID Pasien yang ingin diubah: ")
	fmt.Scan(&id)
	for i := 0; i < m; i++ {
		if a[i].id == id && a[i].ada {
			fmt.Print("Pilih data yang ingin diubah: \n1. Nama\n2. Umur\n3. Paket\n4. Tanggal Registrasi\nPilihan: ")
			fmt.Scan(&pilihan)
			switch pilihan {
			case 1:
				fmt.Print("Masukkan Nama Baru: ")
				fmt.Scan(&nama)
				a[i].nama = nama
			case 2:
				fmt.Print("Masukkan Umur Baru: ")
				fmt.Scan(&umur)
				a[i].umur = umur
			case 3:
				fmt.Print("Masukkan Paket Baru: ")
				fmt.Scan(&paket)
				a[i].paket = paket
			case 4:
				fmt.Print("Masukkan Tanggal Registrasi Baru (yyyy mm dd): ")
				fmt.Scan(&thn, &bln, &tgl)
				a[i].waktu = tanggal{tgl, bln, thn}
			default:
				fmt.Println("Pilihan tidak valid")
			}
			fmt.Println("Data berhasil diubah")
			break
		}
	}
}

func ubahDataPaket(b *tabPaket, p int) { //?procedure untuk mengubah data paket
	var id, biaya, pilihan int
	var nama string

	fmt.Print("Masukkan ID Paket yang ingin diubah: ")
	fmt.Scan(&id)
	for i := 0; i < p; i++ {
		if b[i].id == id && b[i].ada {
			fmt.Print("Pilih data yang ingin diubah: \n1. Nama\n2. Biaya\nPilihan: ")
			fmt.Scan(&pilihan)
			switch pilihan {
			case 1:
				fmt.Print("Masukkan Nama Baru: ")
				fmt.Scan(&nama)
				b[i].nama = nama
			case 2:
				fmt.Print("Masukkan Biaya Baru: ")
				fmt.Scan(&biaya)
				b[i].biaya = biaya
			default:
				fmt.Println("Pilihan tidak valid")
			}
			fmt.Println("Data berhasil diubah")
			break
		}
	}
}

func deleteData(a *tabPasien, b *tabPaket, m, p *int) { //?Procedure untuk menghapus data
	var pilihan int

	fmt.Println("Pilihan Anda untuk menghapus data: ")
	fmt.Println("1. Hapus Data Pasien")
	fmt.Println("2. Hapus Data Paket")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		hapusDataPasien(a, m)
	} else if pilihan == 2 {
		hapusDataPaket(b, p)
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func hapusDataPasien(a *tabPasien, m *int) { //?Procedure untuk menghapus data pasien
	var id int

	fmt.Print("Masukkan ID Pasien yang ingin dihapus: ")
	fmt.Scan(&id)
	for i := 0; i < *m; i++ {
		if a[i].id == id && a[i].ada {
			a[i].ada = false
			fmt.Println("Data Pasien berhasil dihapus")
		}
	}
	fmt.Println("Data Pasien tidak ditemukan")
}

func hapusDataPaket(b *tabPaket, p *int) { //?Procedure untuk menghapus data paket
	var id int

	fmt.Print("Masukkan ID Paket yang ingin dihapus: ")
	fmt.Scan(&id)
	for i := 0; i < *p; i++ {
		if b[i].id == id && b[i].ada {
			b[i].ada = false
			fmt.Println("Data Paket berhasil dihapus")
		}
	}
	fmt.Println("Data Paket tidak ditemukan")
}

func menampilkanData(a *tabPasien, b *tabPaket, m, p *int) {
	var pilihan int

	fmt.Println("Pilihan Anda untuk menampilkan data: ")
	fmt.Println("1. Tampilkan Semua Pasien")
	fmt.Println("2. Tampilkan Semua Paket")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		tampilkanSemuaPasien(a, *m)
	} else if pilihan == 2 {
		tampilkanSemuaPaket(b, *p)
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func tampilkanSemuaPasien(a *tabPasien, m int) { //?Procedure untuk menampilkan seluruh data pasien
	fmt.Println("Daftar Semua Pasien:")
	for i := 0; i < m; i++ {
		if a[i].ada {
			fmt.Printf("ID: %d, Nama: %s, Umur: %d, Paket: %s, Biaya: %d, Tanggal: %02d-%02d-%04d\n",
				a[i].id, a[i].nama, a[i].umur, a[i].paket, a[i].biaya, a[i].waktu.tanggal, a[i].waktu.bulan, a[i].waktu.tahun)
		}
	}
}

func tampilkanSemuaPaket(b *tabPaket, p int) { //?procedure untuk menampilkan seluruh data paket
	fmt.Println("Daftar Semua Paket:")
	for i := 0; i < p; i++ {
		if b[i].ada {
			fmt.Printf("ID: %d, Nama: %s, Biaya: %d\n", b[i].id, b[i].nama, b[i].biaya)
		}
	}
}

func sortingData(a *tabPasien, b *tabPaket, m, p *int) {
	var pilihan int

	fmt.Println("Pilihan Anda untuk mengurutkan data: ")
	fmt.Println("1. Urutkan Pasien berdasarkan Nama")
	fmt.Println("2. Urutkan Paket berdasarkan Nama")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		sortPasienByName(a, *m)
	} else if pilihan == 2 {
		sortPaketByName(b, *p)
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func sortPasienByName(a *tabPasien, m int) {
	sort.SliceStable(a[:m], func(i, j int) bool {
		return a[i].nama < a[j].nama
	})
	fmt.Println("Data Pasien berhasil diurutkan berdasarkan Nama")
}

func sortPaketByName(b *tabPaket, p int) {
	sort.SliceStable(b[:p], func(i, j int) bool {
		return b[i].nama < b[j].nama
	})
	fmt.Println("Data Paket berhasil diurutkan berdasarkan Nama")
}

func Header() {
	fmt.Println("=========== \x1b[32mAvicenna Hospital\x1b[0m ============")
	fmt.Println("=============== \x1b[32mMenu Utama\x1b[0m ===============")
	fmt.Println("|1. Input Data                           |")
	fmt.Println("|2. Ubah Data                            |")
	fmt.Println("|3. Hapus Data                           |")
	fmt.Println("|4. Tampilkan Data                       |")
	fmt.Println("|5. Urutkan Data                         |")
	fmt.Println("|6. Logout                               |")
	fmt.Println("==========================================")
	fmt.Println("|\x1b[31mPilihan anda(1/2/3/4/5/6)?, Kemudian press enter...\x1b[0m ")
}

func HeaderinputData() {
	fmt.Println("=========== Menu Input Data ===========")
	fmt.Println("1. Input Data Pasien")
	fmt.Println("2. Input Data Paket")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Println("4. Keluar")
	fmt.Println("========================================")
}

func Headerinput1() {
	fmt.Println("========== Input Data Pasien ==========")
}

func Headerinput12() {
	fmt.Println("========== Input Data Paket ==========")
}

func headerregistrasi() {
	fmt.Println("========== Registrasi Selesai ==========")
}

func headerregistrasi21(){ // ? Procedure untuk tampilan awal registrasi 

	fmt.Println("========================================================")
	fmt.Println("| 	        Aplikasi Medical Check Up	       |")
	fmt.Println("========================================================")
	fmt.Println("|Silahkan Tentukan Pilihan Anda                        |")
	fmt.Println("|1.📌 Sign Up		                               |")
	fmt.Println("|2.📌 Sign In				               |")
	fmt.Println("|3.❌ Exit			                       |")
	fmt.Println("========================================================")
	fmt.Println("|\x1b[31mPilihan anda(1/2/3)?, Kemudian press enter...\x1b[0m ")		

}
func headerregistrasi2() {
	fmt.Println("========== Registrasi ==========")
	fmt.Println("1. Registrasi")
	fmt.Println("2. Login")
	fmt.Println("3. Keluar")
	fmt.Println("================================")
}

func headerlogin() {
	fmt.Println("========== Login ==========")
	fmt.Println("1. Menu Utama")
	fmt.Println("2. Logout")
	fmt.Println("3. Keluar")
	fmt.Println("================================")
}

func jedawaktu(detik int) { // ? Procedure untuk jeda waktu
	for i := 0; i < (detik)*1000000000; i++ {
	}
}

func cls() { // ? Procedure menghapus atau clear terminal dengan package os
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

	}
}