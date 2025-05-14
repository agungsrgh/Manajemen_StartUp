package data

import (
    "fmt"
    "manajemen-startup/startup"
)

func TampilkanStartup(daftar *[100]startup.Startup, jumlah int) {
    if jumlah == 0 {
        fmt.Println("Belum ada data startup yang dimasukkan.")
        return
    }

    fmt.Println("\n=== Daftar Startup ===")
    for i := 0; i < jumlah; i++ {
        fmt.Printf("%d. Nama: %s\n", i+1, daftar[i].Nama)
        fmt.Printf("   Bidang Usaha: %s\n", daftar[i].Bidang)
        fmt.Printf("   Tahun Berdiri: %d\n", daftar[i].TahunBerdiri)
        fmt.Println("-------------------------------")
    }
}
