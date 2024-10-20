package utils

import (
	"Project/model"
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Input pada field %s tidak boleh kosong", e.Field)
}

func generateNik() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	nik := r.Intn(9000000000000000) + 1000000000000000
	return nik
}

func ValidateName(name string) error {
	re := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	if !re.MatchString(name) {
		return errors.New("Nama harus hanya berisi huruf")
	}
	return nil
}

func validateDate(dateStr string) error {
	_, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		return errors.New("Gunakan format tanggal lahir dd-mm-yyyy")
	}
	return nil
}

func CreateKtp(reader *bufio.Reader) (model.KtpInfo, error) {
	nik := generateNik()

	fmt.Print("Masukkan nama: ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)
	if nama == "" {
		return nil, &ValidationError{"Nama"}
	}

	if err := ValidateName(nama); err != nil {
		return nil, err
	}

	fmt.Print("Masukkan alamat: ")
	alamat, _ := reader.ReadString('\n')
	alamat = strings.TrimSpace(alamat)
	if alamat == "" {
		return nil, &ValidationError{"Alamat"}
	}

	fmt.Print("Masukkan tempat lahir: ")
	tempatLahir, _ := reader.ReadString('\n')
	tempatLahir = strings.TrimSpace(tempatLahir)
	if tempatLahir == "" {
		return nil, &ValidationError{"Tempat Lahir"}
	}
	
	fmt.Print("Masukkan tanggal lahir (dd-mm-yyyy): ")
	tanggalLahir, _ := reader.ReadString('\n')
	tanggalLahir = strings.TrimSpace(tanggalLahir)
	if tanggalLahir == "" {
		return nil, &ValidationError{"Tanggal Lahir"}
	}

	if err := validateDate(tanggalLahir); err != nil {
		return nil, err
	}

	return &model.Ktp{
		Nik: 		  nik,
		Nama: 		  nama,
		Alamat: 	  alamat,
		TempatLahir:  tempatLahir,
		TanggalLahir: tanggalLahir,

	}, nil
}

func PrintKtp(ktp model.KtpInfo) {
	fmt.Println("\n|| ========== Kartu Tanda Penduduk ========== ||")
	fmt.Println()
	fmt.Println("NIK           : ", ktp.GetNik())
	fmt.Println("Nama          : ", ktp.GetNama())
	fmt.Println("Alamat        : ", ktp.GetAlamat())
	fmt.Println("Tempat Lahir  : ", ktp.GetTempatLahir())
	fmt.Println("Tanggal Lahir : ", ktp.GetTanggalLahir())
	fmt.Println()
	fmt.Println("|| =========================================== ||")
}