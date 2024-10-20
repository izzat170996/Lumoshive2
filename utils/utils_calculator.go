package utils

import (
	"Project/model"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func UseCalculator(reader *bufio.Reader) {
	mc := model.MyCalculator{Name: "SimpleCalculator"}
	fmt.Println("\n|| ======= Kalkulator ======= ||")

	for {
		fmt.Println("Pilih operasi:")
		fmt.Println("1. Tambah")
		fmt.Println("2. Kurang")
		fmt.Println("3. Kali")
		fmt.Println("4. Bagi")
		fmt.Println("0. Kembali ke menu utama")

		fmt.Print("Masukkan pilihan: ")
		choiceInput, _ := reader.ReadString('\n')
		choiceInput = strings.TrimSpace(choiceInput)
		choice, err := strconv.Atoi(choiceInput)

		if err != nil || choice < 0 || choice > 4 {
			fmt.Println("Pilihan tidak valid, silakan pilih antara 0-4.")
			continue
		}

		if choice == 0 {
			break
		}

		// Meminta input dua angka dari pengguna
		a, b := GetTwoNumbers(reader)

		var result float64

		// Melakukan operasi sesuai pilihan pengguna
		switch choice {
		case 1:
			result = mc.Add(float64(a), float64(b))
		case 2:
			result = mc.Subtract(float64(a), float64(b))
		case 3:
			result = mc.Multiply(float64(a), float64(b))
		case 4:
			var err error
			result, err = mc.Divide(float64(a), float64(b))
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
		}

		fmt.Printf("Hasil: %s\n", FormatResult(result))
	}
}

func GetTwoNumbers(reader *bufio.Reader) (float64, float64) {
	fmt.Print("Masukkan angka pertama: ")
	aInput, _ := reader.ReadString('\n')
	aInput = strings.TrimSpace(aInput)
	a, _ := strconv.ParseFloat(aInput, 64)

	fmt.Print("Masukkan angka kedua: ")
	bInput, _ := reader.ReadString('\n')
	bInput = strings.TrimSpace(bInput)
	b, _ := strconv.ParseFloat(bInput, 64)

	return a, b
}

func FormatResult(result float64) string {
	if result == float64(int(result)) {
		return fmt.Sprintf("%d", int(result))
	}
	return fmt.Sprintf("%.2f", result)
}
