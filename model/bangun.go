package model

import "errors"

type Rumus interface {
	Luas() (float64, error)
	Keliling() (float64, error)
}

type Persegi struct {
	Sisi float64
}

func (p Persegi) Luas() (float64, error) {
	if p.Sisi <= 0 {
		return 0, errors.New("Nilai sisi tidak boleh 0 atau negatif")
	}
	return p.Sisi * p.Sisi, nil
}

func (p Persegi) Keliling() (float64, error) {
	if p.Sisi <= 0 {
		return 0, errors.New("Nilai sisi tidak boleh 0 atau negatif")
	}
	return 4 * p.Sisi, nil
}

type PersegiPanjang struct {
	Panjang float64
	Lebar  float64
}

func (pp PersegiPanjang) Luas() (float64, error) {
	if pp.Panjang <= 0 || pp.Lebar <= 0 {
		return 0, errors.New("Nilai panjang dan lebar tidak boleh 0 atau negatif")
	}
	return pp.Panjang * pp.Lebar, nil
}

func (pp PersegiPanjang) Keliling() (float64, error) {
	if pp.Panjang <= 0 || pp.Lebar <= 0 {
		return 0, errors.New("Nilai panjang dan lebar tidak boleh 0 atau negatif")
	}
	return 2 * (pp.Panjang + pp.Lebar), nil
}

type Lingkaran struct {
	JariJari float64
}

func (l Lingkaran) Luas() (float64, error) {
	if l.JariJari <= 0 {
		return 0, errors.New("Nilai jari-jari tidak boleh 0 atau negatif")
	}
	return 3.14 * l.JariJari * l.JariJari, nil
}

func (l Lingkaran) Keliling() (float64, error) {
	if l.JariJari <= 0 {
		return 0, errors.New("Nilai jari-jari tidak boleh 0 atau negatif")
	}
	return 2 * 3.14 * l.JariJari, nil
}

type Segitiga struct {
	Alas   float64
	Tinggi float64
}

func (s Segitiga) Luas() (float64, error) {
	if s.Alas <= 0 || s.Tinggi <= 0 {
		return 0, errors.New("Nilai alas dan tinggi tidak boleh 0 atau negatif")
	}
	return 0.5 * s.Alas * s.Tinggi, nil
}
