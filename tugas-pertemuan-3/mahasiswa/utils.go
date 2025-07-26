package mahasiswa

import "fmt"

var maxNilai int = 100
var Versi string = "v1.0.0"

func hitungRataRata(nilai ...int) float64 {
	sum := 0
	for i := range nilai {
		sum += nilai[i]
	}
	result := float64(sum) / float64(len(nilai))
	return result
}

func BuatMahasiswa(nama string, umur int, nilai ...int) *Mahasiswa {
	return &Mahasiswa{
		Nama:     nama,
		Nilai:    nilai,
		umur:     umur,
		nilaiAvg: hitungRataRata(nilai...),
	}

}

func (m Mahasiswa) Info() string {
	return fmt.Sprintf("Nama: %v, Umur: %v", m.Nama, m.umur)
}

func (m Mahasiswa) RataRata() float64 {
	return m.nilaiAvg
}

func (m Mahasiswa) GetUmur() int {
	return m.umur
}

func Describe(d Deskripsi) {
	fmt.Println(d.Info())
	fmt.Println(d.RataRata())
	fmt.Println(d.GetUmur())
}

func GetMaxNilai() int {
	return maxNilai
}

func TotalAge(age ...int) int {
	result := 0
	for _, v := range age {
		result += v
	}
	return result
}
