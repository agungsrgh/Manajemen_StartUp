package main

import (
    "fmt"
    "manajemen-startup/data"
    "manajemen-startup/search"
    "manajemen-startup/startup"
    "strings"
)

func main() {
    var daftar [100]startup.Startup
    var jumlah int
    var pilihan int

    for {
        // Menampilkan menu
        fmt.Println("=== Aplikasi Manajemen Startup ===")
        fmt.Println("1. Lihat semua startup")
        fmt.Println("2. Cari startup berdasarkan nama")
        fmt.Println("3. Cari startup berdasarkan bidang")
        fmt.Println("4. Tambah startup")
        fmt.Println("5. Keluar")
        fmt.Print("Pilih menu: ")
        fmt.Scanln(&pilihan)

        switch pilihan {
        case 1:
            data.TampilkanStartup(&daftar, jumlah)
        case 2:
            var nama string
            fmt.Print("Masukkan nama startup yang dicari: ")
            fmt.Scanln(&nama)
            search.CariStartupByNama(&daftar, jumlah, strings.TrimSpace(nama))
        case 3:
            var bidang string
            fmt.Print("Masukkan bidang usaha yang dicari: ")
            fmt.Scanln(&bidang)
            search.CariStartupByBidang(&daftar, jumlah, strings.TrimSpace(bidang))
        case 4:
            data.TambahStartup(&daftar, &jumlah)
        case 5:
            fmt.Println("Terima kasih, program selesai.")
            return
        default:
            fmt.Println("Pilihan tidak valid.")
        }
    }
}
