package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

// isiArray mengisi array t dengan karakter dari input pengguna
func isiArray(t *tabel, n *int) {
	/* I.S. Data tersedia dalam piranti masukan
	   F.S. Array t berisi sejumlah n karakter yang dimasukkan user,
	   Proses input selama karakter bukanlah TITIK dan n <= NMAX
	*/
	var input string
	fmt.Print("Masukkan teks (akhiri dengan titik '.'): ")
	fmt.Scan(&input)

	for i, char := range input {
		if i >= NMAX || char == '.' {
			break
		}
		t[i] = char
		*n++
	}
}

// cetakArray mencetak n karakter dari array t
func cetakArray(t tabel, n int) {
	/* I.S. Terdefinisi array t yang berisi sejumlah n karakter
	   F.S. n karakter dalam array muncul di layar
	*/
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

// balikanArray membalik urutan karakter dalam array t
func balikanArray(t *tabel, n int) {
	/* I.S. Terdefinisi array t yang berisi sejumlah n karakter
	   F.S. Urutan isi array t terbalik
	*/
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
}

// palindrom memeriksa apakah susunan karakter di dalam t membentuk palindrom
func palindrom(t tabel, n int) bool {
	/* Mengembalikan true apabila susunan karakter di dalam t
	   membentuk palindrom, dan false apabila sebaliknya.
	   Petunjuk: Manfaatkan prosedur balikanArray
	*/
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)
	fmt.Print("Teks: ")
	cetakArray(tab, m)

	if palindrom(tab, m) {
		fmt.Println("Palindrom? true")
	} else {
		fmt.Println("Palindrom? false")
	}

	balikanArray(&tab, m)
	fmt.Print("Reverse teks: ")
	cetakArray(tab, m)
}
