package search

import (
    "fmt"
    "manajemen-startup/startup"
)

func CariStartupByNama(daftar *[100]startup.Startup, jumlah int, nama string) {
    found := false
    for i := 0; i < jumlah; i++ {
        if daftar[i].Nama == nama {
            fmt.Printf("\nDitemukan: %s\n", daftar[i].Nama)
            fmt.Printf("Bidang Usaha: %s\n", daftar[i].Bidang)
            fmt.Printf("Tahun Berdiri: %d\n", daftar[i].TahunBerdiri)
            found = true
            break
        }
    }
    if !found {
        fmt.Println("Startup dengan nama itu tidak ditemukan.")
    }
}

func CariStartupByBidang(daftar *[100]startup.Startup, jumlah int, bidang string) {
    found := false
    fmt.Printf("\nMencari startup di bidang: %s\n", bidang)
    for i := 0; i < jumlah; i++ {
        if daftar[i].Bidang == bidang {
            fmt.Printf("Nama: %s\n", daftar[i].Nama)
            fmt.Printf("Tahun Berdiri: %d\n", daftar[i].TahunBerdiri)
            found = true
        }
    }
    if !found {
        fmt.Println("Tidak ada startup di bidang ini.")
    }
}
