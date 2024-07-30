package main

type Geometri interface {
	luas()
	keliling()
}

type Segitiga struct {
	alas   float32
	tinggi float32
}

type Persegi struct {
	alas uint8
}

func luas(alas, tinggi int) float32 {
	return (float32)(alas) * (float32)(tinggi) / 2
}
func keliling(sisi int) int {
	return sisi + sisi + sisi
}

func (s Segitiga) luas() float32 {
	return s.alas * s.tinggi / 2
}

func (s Segitiga) keliling() float32 {
	return s.alas * 3
}

func (p Persegi) luas() uint8 {
	return p.alas * p.alas
}

func (p Persegi) keliling() uint8 {
	return p.alas * 4
}
