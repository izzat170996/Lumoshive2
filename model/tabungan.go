package model

import "errors"

type TabunganInfo interface {
	GetNama() string
	GetSimpan() (float64, error)
	GetPinjam() (float64, error)
	GetJumlah() float64
}

type Tabungan struct {
	Nama string
	Simpan float64
	Pinjam float64
	Jumlah float64
}

func (t Tabungan) GetNama() string {
	return t.Nama
}

func (t Tabungan) GetSimpan() (float64, error) {
	if t.Simpan <= 0 {
		return 0, errors.New("Menyimpan saldo harus memiliki nominal, tidak boleh 0")
	}
	return t.Simpan, nil
}

func (t Tabungan) GetPinjam() (float64, error) {
	if t.Pinjam > t.Jumlah {
		return 0, errors.New("Saldo yang anda miliki tidak cukup")
	}
	return t.Pinjam, nil
}

func (t Tabungan) GetJumlah() float64 {
	return t.Jumlah
}