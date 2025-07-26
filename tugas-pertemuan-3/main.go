package main

import (
	"fmt"
	"tugas-pertemuan-3/mahasiswa"
)

func Hello(name string) string {
	msg := fmt.Sprintf("Hi, %v. Welcome!", name)
	return msg
}

func main() {
	m1 := mahasiswa.BuatMahasiswa("ragil", 20, 91, 93, 81, 91, 90)
	m2 := mahasiswa.BuatMahasiswa("burhan", 25, 90, 90, 90, 80, 100)

	//
	students := []*mahasiswa.Mahasiswa{m1, m2}

	totalAge := mahasiswa.TotalAge()

	for _, val := range students {
		totalAge(val.GetUmur())
		fmt.Println(val.Info())
		fmt.Printf("Rata-rata nilai: %v\n", val.RataRata())
		fmt.Println("---")
	}

	fmt.Printf("Versi Package: %v\n", mahasiswa.Versi)
	fmt.Printf("Nilai Maksimum: %v\n", mahasiswa.GetMaxNilai())
	fmt.Printf("Total Umur Mahasiswa: %v\n", totalAge(0))

}
