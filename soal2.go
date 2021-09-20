package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Persegi struct {
	Sisi float64
}

func (p *Persegi) Luas() float64 {
	return p.Sisi * p.Sisi
}

func (p *Persegi) Keliling() float64 {
	return 4 * p.Sisi
}

type Segitiga struct {
	Alas   float64
	Tinggi float64
}

func (s *Segitiga) Luas() float64 {
	return 0.5 * s.Alas * s.Tinggi
}

func (s *Segitiga) Keliling() int {
	sisi := math.Pow(s.Alas, 2) + math.Pow(s.Tinggi, 2)
	sisi = math.Sqrt(sisi)
	return int(s.Alas) + int(s.Tinggi) + int(sisi)
}

func menuPilihan() {
	fmt.Println("Program Hitung Bangun")
	fmt.Println("================")
	fmt.Println("1. Segitiga")
	fmt.Println("2. Persegi")
	fmt.Println("0. Keluar")
	fmt.Println("================")
}

func menuSegitiga() {
	var alas, tinggi float64
	var prompt string

	for {
		fmt.Println("Program Hitung Bangun Segitiga")

		fmt.Print("Input Alas Segitiga: ")
		fmt.Scan(&alas)

		fmt.Print("Input Tinggi Segitiga: ")
		fmt.Scan(&tinggi)

		segitiga := Segitiga{
			Alas:   alas,
			Tinggi: tinggi,
		}
		fmt.Println("Hasil hitung:")
		fmt.Printf("Luas: %v\n", segitiga.Luas())
		fmt.Printf("Keliling: %v\n", segitiga.Keliling())

		fmt.Printf("Hitung lagi? (y/t): ")
		fmt.Scan(&prompt)

		if strings.ToLower(prompt) == "y" {
			continue
		} else if strings.ToLower(prompt) == "t" {
			break
		} else {
			fmt.Println("silahkan ketik y/t")
			continue
		}
	}

}

func menuPersegi() {
	var sisi float64
	var prompt string

	for {
		fmt.Println("Program Hitung Bangun Persegi")

		fmt.Print("Input Sisi Persegi: ")
		fmt.Scan(&sisi)

		persegi := Persegi{Sisi: sisi}
		fmt.Println("Hasil hitung:")
		fmt.Printf("Luas: %v\n", persegi.Luas())
		fmt.Printf("Keliling: %v\n", persegi.Keliling())

		fmt.Printf("Hitung lagi? (y/t): ")
		fmt.Scan(&prompt)

		if strings.ToLower(prompt) == "y" {
			continue
		} else if strings.ToLower(prompt) == "t" {
			break
		} else {
			fmt.Println("silahkan ketik y/t")
			continue
		}
	}
}

func main() {
	var choose int
	for {
		menuPilihan()
		fmt.Print("Pilih > ")
		fmt.Scan(&choose)
		switch choose {
		case 1:
			menuSegitiga()
		case 2:
			menuPersegi()
		case 0:
			fmt.Println("Program berakhir")
			os.Exit(0)
		default:
			fmt.Println("Menu pilihan tidak tersedia!")
		}
	}

}
