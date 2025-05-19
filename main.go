package main

import "fmt"

const NMAX int = 100

type tabAkun [NMAX]akun
type tabStartup [NMAX]startup

type akun struct {
	user string
	pw   string
}

type startup struct {
	name, produk, bidang     string
	tahun, dana, jumKaryawan int
	karyawan                 [NMAX]string
	role                     [NMAX]string
}

var n, j int
var A tabAkun
var S tabStartup
var Teknologi, Pendidikan, Keuangan, Kesehatan, Properti, Travel, Logistik int

func main() {
	var opsi int
	for {
		fmt.Println("=======================================")
		fmt.Println("           APLIKASI MANAJEMEN          ")
		fmt.Println("                 STARTUP               ")
		fmt.Println("=======================================")
		fmt.Println()
		fmt.Println("Silahkan memilih opsi di bawah")
		fmt.Println("1.Register")
		fmt.Println("2.Login")
		fmt.Println("3.Exit")
		fmt.Print("Masukkan Opsi: ")
		fmt.Scan(&opsi)

		switch opsi {
		case 1:
			register(&A, &n)
		case 2:
			login(A, n)
		case 3:
			fmt.Println("=======================================")
			fmt.Println("              TERIMA KASIH             ")
			fmt.Println("=======================================")
			fmt.Println()
			return
		default:
			fmt.Println("Opsi tidak valid. Silahkan coba lagi!")
			fmt.Println()
		}
	}
}

func register(A *tabAkun, n *int) {
	fmt.Print("Masukkan Username yang ingin anda gunakan : ")
	fmt.Scan(&A[*n].user)
	for i := 0; i < *n; i++ {
		if A[i].user == A[*n].user {
			fmt.Println("Username sudah digunakan. Silakan gunakan username lain.")
			fmt.Println()
			return
		}
	}

	fmt.Print("Masukkan Password yang ingin anda gunakan :")
	fmt.Scan(&A[*n].pw)
	*n = *n + 1
	fmt.Println()
	fmt.Println("=======================================")
	fmt.Println("              AKUN BERHASIL            ")
	fmt.Println("               DITAMBAHKAN             ")
	fmt.Println("=======================================")
	fmt.Println()
}

func login(A tabAkun, n int) {
	var username, password string
	var i int

	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)
	for i = 0; i < n; i++ {
		if A[i].user == username && A[i].pw == password {
			fmt.Println()
			fmt.Println("=======================================")
			fmt.Println("             LOGIN BERHASIL            ")
			fmt.Println("             SELAMAT DATANG            ")
			fmt.Println("=======================================")
			fmt.Println()
			dashboard()
			return
		}
	}
	fmt.Println()
	fmt.Println("===#       LOGIN GAGAL       #===")
	fmt.Println()
}

func dashboard() {
	var opsi int
	for {
		fmt.Println("=======================================")
		fmt.Println("              SELAMAT DATANG           ")
		fmt.Println("                    DI                 ")
		fmt.Println("                MENU UTAMA             ")
		fmt.Println("=======================================")
		fmt.Println()
		fmt.Println("1. Tambahkan Startup")
		fmt.Println("2. Edit Startup")
		fmt.Println("3. Hapus Startup")
		fmt.Println("4. Search")
		fmt.Println("5. Daftar Startup")
		fmt.Println("6. Laporan")
		fmt.Println("7. Logout")
		fmt.Println()
		fmt.Print("Masukkan Opsi: ")
		fmt.Scan(&opsi)

		switch opsi {
		case 1:
			addStartup(&S, &j)
		case 2:
			editStartup(&S, j)
		case 3:
			deleteStartup(&S, &j)
		case 4:
			search()
		case 5:
			listing()
		case 6:
			laporanbidang(S, j)
		case 7:
			return
		default:
			fmt.Println("===#       OPSI TIDAK VALID       #===")
		}
	}
}

func addStartup(S *tabStartup, j *int) {
	var i int
	var opsi int

	fmt.Print("Nama Startup: ")
	fmt.Scan(&S[*j].name)
	for i = 0; i < *j; i++ {
		if S[i].name == S[*j].name {
			fmt.Println("==##       NAMA STARTUP SUDAH DIGUNAKAN       ##==")
			fmt.Println()
			return
		}
	}

	fmt.Print("Nama Produk: ")
	fmt.Scan(&S[*j].produk)
	fmt.Print("Tahun berdiri: ")
	fmt.Scan(&S[*j].tahun)
	fmt.Print("Jumlah dana: ")
	fmt.Scan(&S[*j].dana)
	fmt.Print("Jumlah karyawan tim: ")
	fmt.Scan(&S[*j].jumKaryawan)
	fmt.Println("Masukkan nama dan role")

	for i = 0; i < S[*j].jumKaryawan; i++ {
		fmt.Scan(&S[*j].karyawan[i])
		fmt.Scan(&S[*j].role[i])
	}
	fmt.Println()
	fmt.Println("Pilih bidang startup")
	fmt.Println("1. Teknologi")
	fmt.Println("2. Pendidikan")
	fmt.Println("3. Keuangan")
	fmt.Println("4. Kesehatan")
	fmt.Println("5. Properti")
	fmt.Println("6. Travel")
	fmt.Println("7. Logistik")
	fmt.Print("Masukkan opsi: ")
	fmt.Scan(&opsi)

	switch opsi {
	case 1:
		S[*j].bidang = "Teknologi"
		Teknologi++
	case 2:
		S[*j].bidang = "Pendidikan"
		Pendidikan++
	case 3:
		S[*j].bidang = "Keuangan"
		Keuangan++
	case 4:
		S[*j].bidang = "Kesehatan"
		Kesehatan++
	case 5:
		S[*j].bidang = "Properti"
		Properti++
	case 6:
		S[*j].bidang = "Travel"
		Travel++
	case 7:
		S[*j].bidang = "Logistik"
		Logistik++
	default:
		fmt.Println("===#       BIDANG TIDAK VALID       #===")
	}

	fmt.Println("=======================================")
	fmt.Println("            STARTUP BERHASIL           ")
	fmt.Println("              DITAMBAHKAN              ")
	fmt.Println("=======================================")
	fmt.Println()
	*j = *j + 1
}

func editStartup(S *tabStartup, j int) {
	var carinama string
	var i int
	var index int
	var found bool = false

	fmt.Println("Masukkan nama startup yang ingin diedit: ")
	fmt.Scan(&carinama)

	for i = 0; i < j; i++ {
		if carinama == S[i].name {
			found = true
			index = i
			break
		}
	}
	if found {
		fmt.Println()
		fmt.Println("=======================================")
		fmt.Println("            STARTUP DITEMUKAN          ")
		fmt.Println("           MASUKKAN DATA BARU!         ")
		fmt.Println("=======================================")
		fmt.Print("Nama Startup: ")
		fmt.Scan(&S[index].name)
		fmt.Print("Nama Produk: ")
		fmt.Scan(&S[index].produk)
		fmt.Print("Tahun berdiri: ")
		fmt.Scan(&S[index].tahun)
		fmt.Print("Jumlah dana: ")
		fmt.Scan(&S[index].dana)
		fmt.Print("Jumlah karyawan tim: ")
		fmt.Scan(&S[index].jumKaryawan)
		fmt.Println("Masukkan nama dan role")

		for i = 0; i < S[index].jumKaryawan; i++ {
			fmt.Scan(&S[index].karyawan[i])
			fmt.Scan(&S[index].role[i])
		}
		fmt.Println()
		fmt.Println("=======================================")
		fmt.Println("             STARTUP BERHASIL          ")
		fmt.Println("                  DIEDIT               ")
		fmt.Println("=======================================")
		fmt.Println()
	} else {
		fmt.Println("===#       STARTUP TIDAK DITEMUKAN       #===")
		fmt.Println()
	}
}

func deleteStartup(S *tabStartup, j *int) {
	var carinama string
	var i int
	var index int
	var found bool = false

	fmt.Println("Masukkan nama startup yang ingin dihapus: ")
	fmt.Scan(&carinama)

	for i = 0; i < *j; i++ {
		if carinama == S[i].name {
			found = true
			index = i
			break
		}
	}
	if found {
		for i = index; i < *j-1; i++ {
			S[i] = S[i+1]
		}
		*j = *j - 1
		switch S[index].bidang {
		case "Teknologi":
			Teknologi--
		case "Pendidikan":
			Pendidikan--
		case "Keuangan":
			Keuangan--
		case "Kesehatan":
			Kesehatan--
		case "Properti":
			Properti--
		case "Travel":
			Travel--
		case "Logistik":
			Logistik--
		}

		fmt.Println()
		fmt.Println("=======================================")
		fmt.Println("            STARTUP BERHASIL           ")
		fmt.Println("                DIHAPUS                ")
		fmt.Println("=======================================")
		fmt.Println()
	} else {
		fmt.Println("===#       STARTUP TIDAK DITEMUKAN       #===")
		fmt.Println()
	}

}

func search() {
	var opsi int

	fmt.Println("Pilih tipe pencarian")
	fmt.Println("1. Pencarian berdasarkan nama")
	fmt.Println("2. Pencarian berdasarkan bidang")
	fmt.Println()
	fmt.Print("Masukkan Opsi: ")
	fmt.Scan(&opsi)
	if opsi == 1 {
		byname(S, j)
	} else if opsi == 2 {
		byfield(S, j)
	}
}

func byname(S tabStartup, j int) {
	var carinama string
	var found bool = false
	var i, k int

	fmt.Println("Masukkan nama startup yang ingin dicari: ")
	fmt.Scan(&carinama)

	for i = 0; i < j; i++ {
		if carinama == S[i].name {
			found = true
			fmt.Println()
			fmt.Println("=======================================")
			fmt.Println("            STARTUP DITEMUKAN          ")
			fmt.Println("=======================================")
			fmt.Println("Nama Startup:", S[i].name)
			fmt.Println("Nama Produk:", S[i].produk)
			fmt.Println("Tahun Berdiri:", S[i].tahun)
			fmt.Println("Jumlah Dana:", S[i].dana)
			fmt.Println("Jumlah Karyawan:", S[i].jumKaryawan)
			fmt.Println("Daftar Karyawan dan Role:")
			for k = 0; k < S[i].jumKaryawan; k++ {
				fmt.Printf("%-10s %-10s", S[i].karyawan[k], S[i].role[k])
				fmt.Println()
			}
			break

		}
	}
	if !found {
		fmt.Println()
		fmt.Println("===#       STARTUP TIDAK DITEMUKAN       #===")
		fmt.Println()
	}
}

func byfield(S tabStartup, j int) {
	var bidang string
	var found bool = false
	var opsi, i int

	fmt.Println("Pilih bidang startup yang ingin dicari:")
	fmt.Println("1. Teknologi")
	fmt.Println("2. Pendidikan")
	fmt.Println("3. Keuangan")
	fmt.Println("4. Kesehatan")
	fmt.Println("5. Properti")
	fmt.Println("6. Travel")
	fmt.Println("7. Logistik")
	fmt.Print("Masukkan pilihan: ")

	fmt.Scan(&opsi)

	if opsi == 1 {
		bidang = "Teknologi"
	} else if opsi == 2 {
		bidang = "Pendidikan"
	} else if opsi == 3 {
		bidang = "Keuangan"
	} else if opsi == 4 {
		bidang = "Kesehatan"
	} else if opsi == 5 {
		bidang = "Properti"
	} else if opsi == 6 {
		bidang = "Travel"
	} else if opsi == 7 {
		bidang = "Logistik"
	}

	fmt.Println()
	fmt.Println("================================")
	fmt.Println("Startup pada bidang", bidang, ":")
	fmt.Println("================================")
	fmt.Println()
	for i = 0; i < j; i++ {
		if bidang == S[i].bidang {
			found = true
			fmt.Println("Nama Startup:", S[i].name)
			fmt.Println("Nama Produk:", S[i].produk)
			fmt.Println("Tahun Berdiri:", S[i].tahun)
			fmt.Println("Jumlah Dana:", S[i].dana)
			fmt.Println()
		}
	}
	if !found {
		fmt.Println()
		fmt.Println("===#       Tidak ada startup dalam bidang tersebut       #===")
		fmt.Println()
	}
}
func listing() {
	var opsi1, opsi2 int

	fmt.Println()
	fmt.Println("Pilih jenis list")
	fmt.Println("1. Berdasarkan Tahun")
	fmt.Println("2. Berdasarkan Pendanaan")
	fmt.Print("Masukkan Pilihan: ")
	fmt.Scan(&opsi1)
	if opsi1 == 1 {
		fmt.Println()
		fmt.Println("Pilih pengurutan")
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&opsi2)
		if opsi2 == 1 {
			fmt.Println()
			listbyyearasc(&S, j)
		}
		if opsi2 == 2 {
			fmt.Println()
			listbyyeardesc(&S, j)
		}
	} else if opsi1 == 2 {
		fmt.Println()
		fmt.Println("Pilih pengurutan")
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&opsi2)
		if opsi2 == 1 {
			fmt.Println()
			listbyfundasc(&S, j)
		}
		if opsi2 == 2 {
			fmt.Println()
			listbyfunddesc(&S, j)
		}
	}
}

func listbyyeardesc(S *tabStartup, j int) {
	var i, idx, pass int
	var temp startup

	pass = 1
	for pass < j {
		idx = pass - 1
		i = pass
		for i < j {
			if S[i].tahun > S[idx].tahun {
				idx = i
			}
			i = i + 1
		}
		temp = S[pass-1]
		S[pass-1] = S[idx]
		S[idx] = temp
		pass = pass + 1
	}
	fmt.Println()
	fmt.Println("=======================================")
	fmt.Println("              DAFTAR STARTUP           ")
	fmt.Println("            TERURUT DESCENDING         ")
	fmt.Println("=======================================")
	fmt.Printf("|%-15s|%-15s|%-10s|\n", "NAMA", "BIDANG", "TAHUN")
	for i = 0; i < j; i++ {
		fmt.Printf("|%-15s|%-15s|%-10d|\n", S[i].name, S[i].bidang, S[i].tahun)
	}
	fmt.Println()
}

func listbyyearasc(S *tabStartup, j int) {
	var i, idx, pass int
	var temp startup

	pass = 1
	for pass < j {
		idx = pass - 1
		i = pass
		for i < j {
			if S[i].tahun < S[idx].tahun {
				idx = i
			}
			i = i + 1
		}
		temp = S[pass-1]
		S[pass-1] = S[idx]
		S[idx] = temp
		pass = pass + 1
	}
	fmt.Println()
	fmt.Println("=======================================")
	fmt.Println("              DAFTAR STARTUP           ")
	fmt.Println("             TERURUT ASCENDING         ")
	fmt.Println("=======================================")
	fmt.Printf("|%-15s|%-15s|%-10s|\n", "NAMA", "BIDANG", "TAHUN")
	for i = 0; i < j; i++ {
		fmt.Printf("|%-15s|%-15s|%-10d|\n", S[i].name, S[i].bidang, S[i].tahun)
	}
	fmt.Println()
}

func listbyfunddesc(S *tabStartup, j int) {
	var i, idx, pass int
	var temp startup

	pass = 1
	for pass < j {
		idx = pass - 1
		i = pass
		for i < j {
			if S[i].dana > S[idx].dana {
				idx = i
			}
			i = i + 1
		}
		temp = S[pass-1]
		S[pass-1] = S[idx]
		S[idx] = temp
		pass = pass + 1
	}
	fmt.Println()
	fmt.Println("=======================================")
	fmt.Println("              DAFTAR STARTUP           ")
	fmt.Println("            TERURUT DESCENDING         ")
	fmt.Println("=======================================")
	fmt.Printf("|%-15s|%-15s|%-10s|\n", "NAMA", "BIDANG", "DANA")
	for i = 0; i < j; i++ {
		fmt.Printf("|%-15s|%-15s|%-10d|\n", S[i].name, S[i].bidang, S[i].dana)
	}
	fmt.Println()
}

func listbyfundasc(S *tabStartup, j int) {
	var i, idx, pass int
	var temp startup

	pass = 1
	for pass < j {
		idx = pass - 1
		i = pass
		for i < j {
			if S[i].dana < S[idx].dana {
				idx = i
			}
			i = i + 1
		}
		temp = S[pass-1]
		S[pass-1] = S[idx]
		S[idx] = temp
		pass = pass + 1
	}
	fmt.Println("=======================================")
	fmt.Println("              DAFTAR STARTUP           ")
	fmt.Println("             TERURUT ASCENDING         ")
	fmt.Println("=======================================")
	fmt.Printf("|%-15s|%-15s|%-10s|\n", "NAMA", "BIDANG", "DANA")
	for i = 0; i < j; i++ {
		fmt.Printf("|%-15s|%-15s|%-10d|\n", S[i].name, S[i].bidang, S[i].dana)
	}
	fmt.Println()
}

func laporanbidang(S tabStartup, j int) {
	var bidang string
	var found bool = false
	var opsi, i int

	fmt.Println("Pilih bidang startup yang ingin dicari:")
	fmt.Println("1. Teknologi")
	fmt.Println("2. Pendidikan")
	fmt.Println("3. Keuangan")
	fmt.Println("4. Kesehatan")
	fmt.Println("5. Properti")
	fmt.Println("6. Travel")
	fmt.Println("7. Logistik")
	fmt.Print("Masukkan pilihan: ")

	fmt.Scan(&opsi)

	if opsi == 1 {
		bidang = "Teknologi"
		fmt.Println("=========================================================")
		fmt.Println("Menampilkan", Teknologi, "startup pada bidang", bidang)
		fmt.Println("=========================================================")
	} else if opsi == 2 {
		bidang = "Pendidikan"
		fmt.Println("=========================================================")
		fmt.Println("Menampilkan", Pendidikan, "startup pada bidang", bidang)
		fmt.Println("=========================================================")
	} else if opsi == 3 {
		bidang = "Keuangan"
		fmt.Println("=========================================================")
		fmt.Println("Menampilkan", Keuangan, "startup pada bidang", bidang)
		fmt.Println("=========================================================")
	} else if opsi == 4 {
		bidang = "Kesehatan"
		fmt.Println("=========================================================")
		fmt.Println("Menampilkan", Kesehatan, "startup pada bidang", bidang)
		fmt.Println("=========================================================")
	} else if opsi == 5 {
		bidang = "Properti"
		fmt.Println("=========================================================")
		fmt.Println("Menampilkan", Properti, "startup pada bidang", bidang)
		fmt.Println("=========================================================")
	} else if opsi == 6 {
		bidang = "Travel"
		fmt.Println("=========================================================")
		fmt.Println("Menampilkan", Travel, "startup pada bidang", bidang)
		fmt.Println("=========================================================")
	} else if opsi == 7 {
		bidang = "Logistik"
		fmt.Println("=========================================================")
		fmt.Println("Menampilkan", Logistik, "startup pada bidang", bidang)
		fmt.Println("=========================================================")
	}

	fmt.Println()
	for i = 0; i < j; i++ {
		if bidang == S[i].bidang {
			found = true
			fmt.Println("Nama Startup:", S[i].name)
			fmt.Println("Nama Produk:", S[i].produk)
			fmt.Println("Tahun Berdiri:", S[i].tahun)
			fmt.Println("Jumlah Dana:", S[i].dana)
			fmt.Println()
		}
	}
	if !found {
		fmt.Println()
		fmt.Println("===#       TIDAK TERDAPAT STARTUP       #===")
	}
}
