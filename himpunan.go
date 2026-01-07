package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Himpunan() {
	rand.Seed(time.Now().UnixNano())

	// Soal 1: Operasi Himpunan Kompleks
	limit := 80
	threshold := limit / 4

	setA := make([]int, 3)
	for i := range setA {
		setA[i] = rand.Intn(limit) + 1
	}

	setB := make([]int, 4)
	for i := range setB {
		setB[i] = rand.Intn(limit) + 1
	}

	setC := make([]int, 2)
	for i := range setC {
		setC[i] = rand.Intn(limit) + 1
	}

	setA = hapusDuplikat(setA)
	setB = hapusDuplikat(setB)
	setC = hapusDuplikat(setC)

	fmt.Printf("\n* Soal 1: Operasi Hitungan KOmpleks")
	fmt.Printf("\nBatas N: %d\n", limit)
	fmt.Printf("A: %v | B: %v | C: %v\n", setA, setB, setC)

	symDiffAB := []int{}
	for _, val := range setA {
		if !adaDiSlice(setB, val) {
			symDiffAB = append(symDiffAB, val)
		}
	}
	for _, val := range setB {
		if !adaDiSlice(setA, val) {
			symDiffAB = append(symDiffAB, val)
		}
	}

	R := []int{}
	for _, val := range symDiffAB {
		if !adaDiSlice(setC, val) {
			R = append(R, val)
		}
	}

	fmt.Printf("Hasil Operasi R: %v\n", R)

	filtered := []int{}
	for _, val := range R {
		if val%2 == 0 && val < threshold {
			filtered = append(filtered, val)
		}
	}

	fmt.Printf("Hasil Filter (Genap < %d): %v\n", threshold, filtered)

	// Soal 2: Analisis Pasangan Subset
	setS := []int{rand.Intn(10) + 1, rand.Intn(10) + 1, rand.Intn(10) + 1, rand.Intn(10) + 1}
	setS = hapusDuplikat(setS)
	K := 10

	fmt.Printf("\n* Soal 2: Analiss Pasangan Subset")
	fmt.Printf("\nSet S: %v | Target K: %d\n", setS, K)
	fmt.Println("Subset 2-Elemen (Sum > K):")

	nomor := 0
	for i := 0; i < len(setS); i++ {
		for j := i + 1; j < len(setS); j++ {
			jumlah := setS[i] + setS[j]
			if jumlah > K {
				nomor = nomor + 1
				fmt.Printf("%d. {%d, %d} (Sum = %d)\n", nomor, setS[i], setS[j], jumlah)
			}
		}
	}

	fmt.Printf("Total: %d Pasangan\n", nomor)
}

func adaDiSlice(slice []int, cari int) bool {
	for _, item := range slice {
		if item == cari {
			return true
		}
	}
	return false
}

func hapusDuplikat(slice []int) []int {
	hasil := make([]int, 0)
	for _, val := range slice {
		if !adaDiSlice(hasil, val) {
			hasil = append(hasil, val)
		}
	}
	return hasil
}
