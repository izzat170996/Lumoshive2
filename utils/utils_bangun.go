package utils

import (
	"Project/model"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func formatResult(result float64) string {
	if result == float64(int(result)) {
		return fmt.Sprintf("%d", int(result))
	}

	return fmt.Sprintf("%.2f", result)
}

func UseRumusBangun(reader *bufio.Reader) {
	fmt.Println("\n|| ======= Rumus Bangun ======= ||")
	fmt.Println("Pilih bangun:")
	fmt.Println("1. Persegi")
	fmt.Println("2. Persegi Panjang")
	fmt.Println("3. Lingkaran")
	fmt.Println("4. Segitiga")
	fmt.Println("0. Kembali ke menu utama") 

	fmt.Print("Masukkan pilihan: ")
	choiceInput, _ := reader.ReadString('\n')
	choiceInput = strings.TrimSpace(choiceInput)
	choice, err := strconv.Atoi(choiceInput)

	if err != nil || choice < 0 || choice > 4 {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	if choice == 0 {
		return
	}

	switch choice {
	case 1:
		// Input untuk persegi
		fmt.Print("Masukkan panjang sisi persegi: ")
		sisiInput, _ := reader.ReadString('\n')
		sisiInput = strings.TrimSpace(sisiInput)
		sisi, err := strconv.ParseFloat(sisiInput, 64)
		if err != nil || sisi <= 0 {
			fmt.Println("Input tidak valid, sisi harus angka positif.")
			return
		}

		p := model.Persegi{Sisi: sisi}
		luas, err := p.Luas()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		keliling, err := p.Keliling()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Luas Persegi: %s\n", formatResult(luas))
		fmt.Printf("Keliling Persegi: %s\n", formatResult(keliling))

	case 2:
		// Input untuk persegi panjang
		fmt.Print("Masukkan panjang persegi panjang: ")
		panjangInput, _ := reader.ReadString('\n')
		panjangInput = strings.TrimSpace(panjangInput)
		panjang, err := strconv.ParseFloat(panjangInput, 64)
		if err != nil || panjang <= 0 {
			fmt.Println("Input tidak valid, panjang harus angka positif.")
			return
		}

		fmt.Print("Masukkan lebar persegi panjang: ")
		lebarInput, _ := reader.ReadString('\n')
		lebarInput = strings.TrimSpace(lebarInput)
		lebar, err := strconv.ParseFloat(lebarInput, 64)
		if err != nil || lebar <= 0 {
			fmt.Println("Input tidak valid, lebar harus angka positif.")
			return
		}

		pp := model.PersegiPanjang{Panjang: panjang, Lebar: lebar}
		luas, err := pp.Luas()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		keliling, err := pp.Keliling()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Luas Persegi Panjang: %s\n", formatResult(luas))
		fmt.Printf("Keliling Persegi Panjang: %s\n", formatResult(keliling))

	case 3:
		// Input untuk lingkaran
		fmt.Print("Masukkan jari-jari lingkaran: ")
		jariJariInput, _ := reader.ReadString('\n')
		jariJariInput = strings.TrimSpace(jariJariInput)
		jariJari, err := strconv.ParseFloat(jariJariInput, 64)
		if err != nil || jariJari <= 0 {
			fmt.Println("Input tidak valid, jari-jari harus angka positif.")
			return
		}

		l := model.Lingkaran{JariJari: jariJari}
		luas , err := l.Luas()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		keliling, err := l.Keliling()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Luas Lingkaran: %s\n", formatResult(luas))
		fmt.Printf("Keliling Lingkaran: %s\n", formatResult(keliling))

	case 4:
		// Input untuk segitiga
		fmt.Print("Masukkan alas segitiga: ")
		alasInput, _ := reader.ReadString('\n')
		alasInput = strings.TrimSpace(alasInput)
		alas, err := strconv.ParseFloat(alasInput, 64)
		if err != nil || alas <= 0 {
			fmt.Println("Input tidak valid, alas harus angka positif.")
			return
		}

		fmt.Print("Masukkan tinggi segitiga: ")
		tinggiInput, _ := reader.ReadString('\n')
		tinggiInput = strings.TrimSpace(tinggiInput)
		tinggi, err := strconv.ParseFloat(tinggiInput, 64)
		if err != nil || tinggi <= 0 {
			fmt.Println("Input tidak valid, tinggi harus angka positif.")
			return
		}

		s := model.Segitiga{Alas: alas, Tinggi: tinggi}
		luas, err := s.Luas()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Luas Segitiga: %s\n", formatResult(luas))
		
	}
}
