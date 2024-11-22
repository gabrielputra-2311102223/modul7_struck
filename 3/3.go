package main

import (
	"fmt"
)

func main() {
	var clubA, clubB string
	var scoreA, scoreB int

	// Input nama klub
	fmt.Print("Klub A: ")
	fmt.Scanln(&clubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&clubB)

	// Slice untuk menyimpan pemenang
	var winners []string

	// Loop untuk memasukkan skor pertandingan
	for i := 1; ; i++ {
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scanln(&scoreA, &scoreB)

		// Menghentikan input jika skor negatif
		if scoreA < 0 || scoreB < 0 {
			break
		}

		// Menentukan pemenang
		if scoreA > scoreB {
			winners = append(winners, clubA)
		} else if scoreB > scoreA {
			winners = append(winners, clubB)
		} else {
			winners = append(winners, "Draw")
		}
	}

	// Menampilkan hasil pertandingan
	fmt.Println("Hasil pertandingan:")
	for i, winner := range winners {
		fmt.Printf("Hasil %d: %s\n", i+1, winner)
	}

	fmt.Println("Pertandingan selesai")
}
