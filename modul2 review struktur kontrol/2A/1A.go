package main

import "fmt"

func main() {
	var (
		satu, dua, tiga, temp string
	)

	// Meminta input dari pengguna
	fmt.Print("Masukkan input string pertama: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukkan input string kedua: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukkan input string ketiga: ")
	fmt.Scanln(&tiga)

	// Menampilkan output awal
	fmt.Println("Output awal =", satu, dua, tiga)

	// Proses penukaran nilai
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	// Menampilkan output akhir setelah penukaran
	fmt.Println("Output akhir =", satu, dua, tiga)
}
