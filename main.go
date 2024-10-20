package main

import (
	"Project/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	
)

func main() {
	fmt.Println("|| ========== Selamat Datang di Aplikasi Command Line Interface ========== ||")
	reader := bufio.NewReader(os.Stdin)

	for {
		// Menampilkan pilihan aplikasi
		fmt.Println("\nPilih Aplikasi:")
		fmt.Println("1. Kalkulator")
		fmt.Println("2. Rumus Bangun")
		fmt.Println("3. Pembuatan KTP")
		fmt.Println("4. Pembuatan Passport")
		fmt.Println("5. Tabungan")
		fmt.Println("0. Keluar")
		fmt.Print("Masukkan pilihan Anda: ")

		choiceInput, _ := reader.ReadString('\n')
		choiceInput = strings.TrimSpace(choiceInput)
		choice, err := strconv.Atoi(choiceInput)

		if err != nil || choice < 0 || choice > 5 {
			fmt.Println("Pilihan tidak valid, silakan pilih antara 0-5.")
			continue
		}

		if choice == 0 {
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			break
		}

		switch choice {
		case 1:
			utils.UseCalculator(reader)
		case 2:
			utils.UseRumusBangun(reader)
		case 3:
			ktp, err := utils.CreateKtp(reader)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println("\n KTP berhasil dibuat ")
			utils.PrintKtp(ktp)
		case 4:
			passport, err := utils.CreatePassport(reader)
			if err != nil {
				fmt.Println("Error : ", err)
				continue
			}
			fmt.Println("\n Passport berhasil di buat")
			utils.PrintPassport(passport)
		case 5:
			utils.MyTabungan(reader)
		}
		  
	}

	
}