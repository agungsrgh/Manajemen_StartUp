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
	// Menampilkan menu registrasi dan login aplikasi
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
	//I.S. Pengguna memasukkan username dan password yang ingin didaftarkan
	//F.S. Akun berhasil ditambahkan ke dalam sistem
	var duplikat bool
	var i int

	for {
		duplikat = false
		fmt.Print("Masukkan Username yang ingin anda gunakan: ")
		fmt.Scan(&A[*n].user)

		for i = 0; i < *n; i++ {
			if A[i].user == A[*n].user {
				duplikat = true
				break
			}
		}
		if !duplikat {
			break
		}
		fmt.Println("Username sudah digunakan. Silakan gunakan username lain.")
		fmt.Println()
	}

	fmt.Print("Masukkan Password yang ingin anda gunakan: ")
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
	//I.S. Pengguna memasukkan username dan password yang telah didaftarkan
	//F.S. Pengguna berhasil login dan beralih ke menu utama
	var username, password string
	var i int
	var loginBerhasil bool = false

	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)
	for i = 0; i < n; i++ {
		if A[i].user == username && A[i].pw == password {
			loginBerhasil = true
			break
		}
	}
	if loginBerhasil {
		fmt.Println()
		fmt.Println("=======================================")
		fmt.Println("             LOGIN BERHASIL            ")
		fmt.Println("             SELAMAT DATANG            ")
		fmt.Println("=======================================")
		fmt.Println()
		dashboard()
	} else {
		fmt.Println()
		fmt.Println("===#       LOGIN GAGAL       #===")
		fmt.Println()
	}
}

func dashboard() {
	// Menampilkan menu utama dan pilihan menu
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
	//I.S. Pengguna menambahkan Data startup yang ingin didaftarkan
	//F.S. Startup berhasil ditambahkan ke dalam sistem
	var i int
	var opsi int
	var duplikat bool

	for {
		duplikat = false
		fmt.Print("Nama Startup: ")
		fmt.Scan(&S[*j].name)
		for i = 0; i < *j; i++ {
			if S[i].name == S[*j].name {
				duplikat = true
				break
			}
		}
		if !duplikat {
			break
		}
		fmt.Println("==##       NAMA STARTUP SUDAH DIGUNAKAN       ##==")
		fmt.Println()
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
		fmt.Printf("Masukkan nama karyawan ke-%d: ", i+1)
		fmt.Scan(&S[*j].karyawan[i])

		fmt.Printf("Masukkan role karyawan ke-%d: ", i+1)
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
	//I.S. Pengguna mengubah data startup yang telah didaftarkan
	//F.S. Startup yang sudah ada diedit dengan data lain
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
	//I.S. Pengguna memilih startup yang akan di hapus
	//F.S. Startup yang dipilih dihapus dari sistem
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
	// Menampilkan menu pencarian
	var opsi int

	fmt.Println("Pilih tipe pencarian")
	fmt.Println("1. Pencarian berdasarkan nama")
	fmt.Println("2. Pencarian berdasarkan bidang")
	fmt.Println("3. Pencarian berdasarkan tahun berdiri")
	fmt.Print("Masukkan Opsi: ")
	fmt.Scan(&opsi)
	if opsi == 1 {
		byname(S, j)
	} else if opsi == 2 {
		byfield(S, j)
	} else if opsi == 3 {
		var tahun int
		fmt.Print("Masukkan tahun berdiri yang ingin dicari: ")
		fmt.Scan(&tahun)
		binarySearchByYear(S, j, tahun)
	}
}

func byname(S tabStartup, j int) {
	//I.S. Pengguna memasukan nama startup yang ingin dicari
	//F.S. Sistem memberikan tampilan data startup yang dicari
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
			fmt.Println("Daftar Nama Karyawan dan Role:")
			fmt.Printf("%-10s %-15s\n", "Nama", "Role")
			fmt.Println("---------------------------------------")
			for k = 0; k < S[i].jumKaryawan; k++ {
				fmt.Printf("%-10s %-15s", S[i].karyawan[k], S[i].role[k])
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
	//I.S. Pengguna memilih bidang startup yang ingin dicari
	//F.S. Sistem menampilkan startup pada bidang tersebut
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
	found = false
	for i = 0; i < j; i++ {
		if S[i].bidang == bidang {
			found = true
			fmt.Println()
			fmt.Println("Nama Startup :", S[i].name)
			fmt.Println("Nama Produk  :", S[i].produk)
			fmt.Println("Tahun Berdiri:", S[i].tahun)
			fmt.Println("Jumlah Dana  :", S[i].dana)
			fmt.Println("Jumlah Karyawan:", S[i].jumKaryawan)
			fmt.Println("Daftar Nama Karyawan dan Role:")
			fmt.Printf("%-10s %-15s\n", "Nama", "Role")
			fmt.Println("---------------------------------------")
			for k := 0; k < S[i].jumKaryawan; k++ {
				fmt.Printf("%-10s %-15s\n", S[i].karyawan[k], S[i].role[k])
			}
			fmt.Println("---------------------------------------")
		}
	}
	if !found {
		fmt.Println()
		fmt.Println("===#       TIDAK ADA STARTUP DIBIDANG INI       #===")
		fmt.Println()
	}
}

func sortByYearAsc(S *tabStartup, j int) {
	//I.S. Menampilkan daftar startup
	//F.S. Menampilkan daftar startup yang sudah diurutkan berdasarkan tahun secara Selection sort ASCENDING
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
			i++
		}
		temp = S[pass-1]
		S[pass-1] = S[idx]
		S[idx] = temp
		pass++
	}
}

func binarySearchByYear(S tabStartup, j int, target int) {
	// I.S. Pengguna memasukkan tahun  startup yang ingin dicari
	// F.S. Sistem menampilkan startup yang dicari berdasarkan tahun menggunakan metode binary search
	sortByYearAsc(&S, j)
	low := 0
	high := j - 1
	var mid int
	found := false

	for low <= high {
		mid = (low + high) / 2
		if S[mid].tahun == target {
			found = true
			break
		} else if target < S[mid].tahun {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if found {
		fmt.Println()
		fmt.Println("=======================================")
		fmt.Println("         STARTUP DENGAN TAHUN", target)
		fmt.Println("=======================================")

		for i := mid; i >= 0 && S[i].tahun == target; i-- {
			tampilStartup(S[i])
		}
		for i := mid + 1; i < j && S[i].tahun == target; i++ {
			tampilStartup(S[i])
		}
	} else {
		fmt.Println()
		fmt.Println("===#     TIDAK ADA STARTUP DENGAN TAHUN INI     #===")
	}
}

func tampilStartup(s startup) {
	// Layout untuk tampilan data startup yang dicari
	fmt.Println("Nama Startup :", s.name)
	fmt.Println("Produk       :", s.produk)
	fmt.Println("Tahun Berdiri:", s.tahun)
	fmt.Println("Dana         :", s.dana)
	fmt.Println("Bidang       :", s.bidang)
	fmt.Println("Jumlah Karyawan:", s.jumKaryawan)
	fmt.Println("Daftar Karyawan:")
	fmt.Printf("%-10s %-15s\n", "Nama", "Role")
	fmt.Println("---------------------------------------")
	for i := 0; i < s.jumKaryawan; i++ {
		fmt.Printf("%-10s %-15s\n", s.karyawan[i], s.role[i])
	}
	fmt.Println()
}

func listing() {
	// Menampilkan menu daftar startup beserta pilihan menu
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
	//I.S. Menampilkan daftar startup
	//F.S. Menampilkan daftar startup yang telah diurutkan berdasarkan tahun secara Selection sort DESCENDING
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
	//I.S. Menampilkan daftar startup
	//F.S. Menampilkan daftar startup yang telah diurutkan berdasarkan tahun secara Selection sort ASCENDING
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
	//I.S. Menampilkan daftar startup
	//F.S. Menampilkan daftar startup yang telah diurutkan berdasarkan pendanaan secara Insertion sort DESCENDING
	var i, j2 int
	var temp startup

	for i = 1; i < j; i++ {
		temp = S[i]
		j2 = i - 1
		for j2 >= 0 && S[j2].dana < temp.dana {
			S[j2+1] = S[j2]
			j2 = j2 - 1
		}
		S[j2+1] = temp
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
	//I.S. Menampilkan daftar startup
	//F.S. Menampilkan daftar startup yang telah diurutkan berdasarkan pendanaan secara Insertion sort ASCENDING
	var i, j2 int
	var temp startup

	for i = 1; i < j; i++ {
		temp = S[i]
		j2 = i - 1
		for j2 >= 0 && S[j2].dana > temp.dana {
			S[j2+1] = S[j2]
			j2 = j2 - 1
		}
		S[j2+1] = temp
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
	//I.S. Pengguna memilih bidang yang ingin ditampilkan
	//F.S. Sistem menampilkan data startup dari bidang yang dipilih oleh pengguna
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
