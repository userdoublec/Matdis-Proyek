package main

import (
	"fmt"
	"math/rand"
	"time"
)

func MAtrix() {
	rand.Seed(time.Now().UnixNano())

	// Soal 3 : Perkalian dan Trace Matriks
	// la : limit angka
	// N : ordo

	ordo := 3
	la := 10

	fmt.Printf("\n* Soal 3: Perkaliam dan TRace Matriks\n")
	fmt.Printf("Matriks generated (%dx%d)...\n", ordo, ordo)

	matrixA := make([][]int, ordo)
	matrixB := make([][]int, ordo)

	for i := 0; i < ordo; i++ {
		matrixA[i] = make([]int, ordo)
		matrixB[i] = make([]int, ordo)
		for j := 0; j < ordo; j++ {
			matrixA[i][j] = rand.Intn(la) + 1
			matrixB[i][j] = rand.Intn(la) + 1
		}
	}

	fmt.Printf("Matriks A: %v\n", matrixA)
	fmt.Printf("Matriks B: %v\n", matrixB)

	matrixR := make([][]int, ordo)
	for i := 0; i < ordo; i++ {
		matrixR[i] = make([]int, ordo)
		for j := 0; j < ordo; j++ {
			for k := 0; k < ordo; k++ {
				matrixR[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}

	fmt.Printf("Hasil Matriks R : %v\n", matrixR)

	trace := 0
	for i := 0; i < ordo; i++ {
		trace += matrixR[i][i]
	}

	fmt.Printf("Trace: %d\n", trace)

	// Soal 4 : Transformasi Baris
	fmt.Printf("\n* Soal 4: Transformasi Baris\n")

	ordo = 5
	la4 := 15

	matrixM := make([][]int, ordo)
	for i := range matrixM {
		matrixM[i] = make([]int, ordo)
		for j := 0; j < ordo; j++ {
			matrixM[i][j] = rand.Intn(la4) + 1
		}
	}

	// matmax : Matriks Maksimum
	matmax := matrixM[0][0]
	barmax, komax := 0, 0
	// barmax : Baris Maksimum
	// komax : Kolom Maksimum

	for i := 0; i < ordo; i++ {
		for j := 0; j < ordo; j++ {
			if matrixM[i][j] > matmax {
				matmax = matrixM[i][j]
				barmax = i
				komax = j
			}
		}
	}

	fmt.Print("Matrix M (Generated): [")
	for i, row := range matrixM {
		fmt.Printf("%v", row)
		if i < len(matrixM)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")

	fmt.Printf("Menukar Baris 0 dan %d..\n", ordo-1)
	matrixM[0], matrixM[ordo-1] = matrixM[ordo-1], matrixM[0]
	fmt.Printf("Matriks M Terkini : %v\n", matrixM)
	fmt.Printf("Nilai Maksimum : %d ditemukan di Posisi (%d, %d)\n", matmax, barmax, komax)
}
