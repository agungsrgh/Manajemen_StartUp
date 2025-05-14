package data

import (
    "bufio"
    "fmt"
    "manajemen-startup/startup"
    "os"
    "strconv"
    "strings"
)

func TambahStartup(daftar *[100]startup.Startup, jumlah *int) {
    if *jumlah >= 100 {
        fmt.Println("Data startup sudah penuh!")
        return
    }

    reader := bufio.NewReader(os.Stdin)
    var s startup.Startup

    fmt.Print("Nama Startup: ")
    s.Nama, _ = reader.ReadString('\n')
    s.Nama = strings.TrimSpace(s.Nama)

    fmt.Print("Bidang Usaha: ")
    s.Bidang, _ = reader.ReadString('\n')
    s.Bidang = strings.TrimSpace(s.Bidang)

    fmt.Print("Tahun Berdiri: ")
    tahunStr, _ := reader.ReadString('\n')
    tahunStr = strings.TrimSpace(tahunStr)
    s.TahunBerdiri, _ = strconv.Atoi(tahunStr)

    (*daftar)[*jumlah] = s
    *jumlah++

    fmt.Println("Startup berhasil ditambahkan!\n")
}
