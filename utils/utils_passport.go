package utils

import (
	"Project/model"
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
	"golang.org/x/exp/rand"
)

func GenerateULID() string {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(uint64(time.Now().UnixNano()))), 0)
	ulid := ulid.MustNew(ulid.Timestamp(time.Now()), entropy)
	return ulid.String()
}

func ValidateNik(nik string) (int, error) {
	nik = strings.TrimSpace(nik)
	if len(nik) != 16 {
		return 0, errors.New("Nik harus terdiri dari 16 digit")
	}

	nikInt, err := strconv.Atoi(nik)
	if err != nil {
		return 0, errors.New("Nik harus berupa angka")
	}
	return nikInt, nil
}

func CreatePassport(reader *bufio.Reader) (model.PassportInfo, error) {
	noPassport := GenerateULID()

	fmt.Print("Masukan NIK : ")
	nikInput, _ := reader.ReadString('\n')
	nik, err := ValidateNik(nikInput)
	if err != nil {
		return nil, err
	}

	fmt.Print("Masukan Nama : ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)
	if nama == "" {
		return nil, errors.New("Nama tidak boleh kosong")
	}

	if err := ValidateName(nama); err != nil {
		return nil, err
	}

	fmt.Print("Masukan Alamat : ")
	alamat, _ := reader.ReadString('\n')
	alamat = strings.TrimSpace(alamat)
	if alamat == "" {
		return nil, errors.New("Alamat tidak boleh kosong")
	}

	return &model.Passport{
		NoPassport: noPassport,
		Nik: nik,
		Nama: nama,
		Alamat: alamat,
	}, nil
}

func PrintPassport(passport model.PassportInfo) {
	fmt.Println("\n|| ================== Passport ==================  ||")
	fmt.Println()
	fmt.Println("No Passport : ", passport.GetNoPassport())
	fmt.Println("Nik         : ", passport.GetNik())
	fmt.Println("Nama        : ", passport.GetNama())
	fmt.Println("Alamat      : ", passport.GetAlamat())
	fmt.Println()
	fmt.Println("|| ================================================ ||")
}