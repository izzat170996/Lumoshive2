package utils

import (
	"Project/model"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var tabunganSlice []model.Tabungan

func MyTabungan(reader *bufio.Reader) {
    for {
        fmt.Println(" || =============== Selamat Datang Di Tabunganku =============== ||")
        fmt.Println(" || Menu :                                                       ||")
        fmt.Println(" || 1. Menabung                                                  ||")
        fmt.Println(" || 2. Meminjam                                                  ||")
        fmt.Println(" || 3. Cek Tabungan                                              ||")
		fmt.Println(" || 4. Cek User                                                  ||")
        fmt.Println(" || 0. Kembali Ke Menu Utama                                     ||")
        fmt.Println(" || ============================================================ ||")
        fmt.Print("Masukan menu: ")

        choiceInput, _ := reader.ReadString('\n')
        choiceInput = strings.TrimSpace(choiceInput)
        choice, err := strconv.Atoi(choiceInput)

        if err != nil || choice < 0 || choice > 4 {
            fmt.Println("Pilihan tidak valid")
            continue
        }

        if choice == 0 {
            return
        }

        switch choice {
        case 1:
            Menabung(reader)
        case 2:
            if len(tabunganSlice) == 0 {
                fmt.Println("Anda belum menabung, silakan menabung terlebih dahulu!")
                continue
            }
            Meminjam(reader)
        case 3:
            if len(tabunganSlice) == 0 {
                fmt.Println("Anda belum menabung, silakan menabung terlebih dahulu!")
                continue
            }
            CekTabungan(reader)
		case 4:
			if len(tabunganSlice) == 0 {
				fmt.Println("Belum ada user yang menabung")
				continue
			}
			CekUser()
        }
    }
}

func Menabung(reader *bufio.Reader) {
	fmt.Print("Masukan nama : ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	for _, tabungan := range tabunganSlice {
		if tabungan.Nama == nama {
			fmt.Println("Nama sudah terdaftar, gunakan nama lain")
			return
		}
	}

	fmt.Print("Masukan nominal yang ingin ditabung : ")
	simpananInput, _ := reader.ReadString('\n')
	simpananInput = strings.TrimSpace(simpananInput)
	simpan, err := strconv.ParseFloat(simpananInput, 64)
	if err != nil || simpan <= 0 {
		fmt.Println("Nominal tidak boleh 0 atau negatif")
		return
	}

	tabungan := model.Tabungan{
		Nama: nama,
		Simpan: simpan,
		Jumlah: simpan,
	}

	tabunganSlice = append(tabunganSlice, tabungan)
	fmt.Println("Menabung berhasil")
	fmt.Println("Kamu menabung sebesar Rp.", simpan)
}

func Meminjam(reader *bufio.Reader) {
	fmt.Print("Masukan nama yang akan meminjam : ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	var tabungan *model.Tabungan
	for i := range tabunganSlice {
		if tabunganSlice[i].Nama == nama {
			tabungan = &tabunganSlice[i]
			break
		}
	}

	if tabungan == nil {
		fmt.Println("Nama tidak ditemukan")
		return
	}

	fmt.Print("Masukan Jumlah pinjaman : ")
	pinjamInput, _ := reader.ReadString('\n')
	pinjamInput = strings.TrimSpace(pinjamInput)
	pinjam, err := strconv.ParseFloat(pinjamInput, 64)
	if err != nil || pinjam <= 0 {
		fmt.Println("Nominal tidak boleh 0 atau negatif")
		return
	}

	if pinjam > tabungan.Jumlah {
		fmt.Println("Saldo tidak cukup untuk meminjam")
		return
	}

	tabungan.Pinjam = pinjam
	tabungan.Jumlah -= pinjam
	tabungan.Simpan -= pinjam
	
	fmt.Println("Pinjaman berhasil")
}

func CekTabungan(reader *bufio.Reader) {
	fmt.Print("Masukan nama untuk cek tabungan : ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	var tabungan *model.Tabungan
	for _, t := range tabunganSlice {
		if t.Nama == nama {
			tabungan = &t
			break
		}
	}

	if tabungan == nil {
		fmt.Println("Nama tidak ditemukan")
		return
	}

	
	fmt.Println("Nama : ", tabungan.Nama)
	fmt.Println("Saldo : Rp.", tabungan.Jumlah)
	fmt.Println("Saldo yang bisa diambil : Rp.", tabungan.Simpan)
	fmt.Println("Saldo yang harus dibayar : Rp.", tabungan.Pinjam)
}

func CekUser() {
	fmt.Println("Daftar User yang menabung : ")
	for _, tabungan := range tabunganSlice {
		fmt.Printf("Nama : %s, Saldo : Rp. %.2f\n", tabungan.Nama, tabungan.Jumlah)
	}
}