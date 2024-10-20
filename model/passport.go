package model

type PassportInfo interface {
	GetNoPassport() string
	GetNik() int
	GetNama() string
	GetAlamat() string
	GetStatusNegara() string
}

type Passport struct {
	NoPassport string
	Nik        int
	Nama       string
	Alamat     string
}

// GetStatusNegara implements PassportInfo.
func (p *Passport) GetStatusNegara() string {
	panic("unimplemented")
}

func (p Passport) GetNoPassport() string {
	return p.NoPassport
}

func (p Passport) GetNik() int {
	return p.Nik
}

func (p Passport) GetNama() string {
	return p.Nama
}

func (p Passport) GetAlamat() string {
	return p.Alamat
}
