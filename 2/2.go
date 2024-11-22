package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	// Membuat array dengan ukuran n
	array := make([]int, n)
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen %d: ", i)
		fmt.Scan(&array[i])
	}

	// a. Menampilkan isi keseluruhan array
	fmt.Println("a. Isi keseluruhan array:", array)

	// b. Menampilkan elemen-elemen array dengan indeks ganjil
	fmt.Print("b. Elemen-elemen array dengan indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	// c. Menampilkan elemen-elemen array dengan indeks genap
	fmt.Print("c. Elemen-elemen array dengan indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	// d. Menampilkan elemen-elemen array dengan indeks kelipatan x
	var x int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Printf("d. Elemen-elemen array dengan indeks kelipatan %d: ", x)
	for i := 0; i < n; i += x {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	// e. Menghapus elemen pada indeks yang diberikan
	var index int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&index)
	if index >= 0 && index < n {
		array = append(array[:index], array[index+1:]...)
		fmt.Println("e. Isi array setelah penghapusan:", array)
	} else {
		fmt.Println("Indeks tidak valid.")
	}

	// f. Menghitung rata-rata dari elemen dalam array
	sum := 0
	for _, value := range array {
		sum += value
	}
	average := float64(sum) / float64(len(array))
	fmt.Printf("f. Rata-rata dari bilangan dalam array: %.2f\n", average)

	// g. Menghitung standar deviasi dari elemen dalam array
	varianceSum := 0.0
	for _, value := range array {
		varianceSum += math.Pow(float64(value)-average, 2)
	}
	variance := varianceSum / float64(len(array))
	stdDev := math.Sqrt(variance)
	fmt.Printf("g. Standar deviasi dari bilangan dalam array: %.2f\n", stdDev)

	// h. Mencari frekuensi dari bilangan yang diberikan
	var target int
	fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&target)
	frequency := 0
	for _, value := range array {
		if value == target {
			frequency++
		}
	}
	fmt.Printf("h. Frekuensi dari bilangan %d dalam array: %d\n", target, frequency)
}
