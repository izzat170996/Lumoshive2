package model

type KtpInfo interface {
	GetNik() int
	GetNama() string
	GetAlamat() string
	GetTempatLahir() string
	GetTanggalLahir() string
}

type Ktp struct {
	Nik int
	Nama string
	Alamat string
	TempatLahir string
	TanggalLahir string
}

func (k Ktp) GetNik() int {
	return k.Nik
}

func (k Ktp) GetNama() string {
	return k.Nama
}

func (k Ktp) GetAlamat() string {
	return k.Alamat
}

func (k Ktp) GetTempatLahir() string {
	return k.TempatLahir
}

func (k Ktp) GetTanggalLahir() string {
	return k.TanggalLahir
}