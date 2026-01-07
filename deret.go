package main

import (
	"fmt"
)

func Deret() {

	// Soal 5 : Relasi Rekurens Iteratif
	c1, c2, n := 5, -2, 8
	a0, a1 := 1, 1

	fmt.Printf("\n* Soal 5: RElasi Rekursif Iteratif\n")
	fmt.Printf("INPUT: C1 = %d, C2 = %d, N = %d\n", c1, c2, n)
	fmt.Printf("Proses Perhitungan:\n")
	fmt.Printf("Suku 0: %d | Suku 1: %d", a0, a1)

	an := 0
	for i := 2; i <= n; i++ {
		an = (c1 * a1) + (c2 * a0)
		fmt.Printf(" | Suku %d: %d", i, an)

		// Tukar nilainya
		a0 = a1
		a1 = an
	}
	fmt.Printf("\nHASIL AKHIR Suku ke-%d: %d\n", n, an)

	// Soal 6 : Analisis Kedekatan Deret Geometri
	a := 2.0
	r := 0.8
	n6 := 12

	// sth : Sum Tak Hingga
	// sn : Sum Suku Ke-n
	sth := a / (1 - r)

	sn := 0.0
	suku := a

	for i := 1; i <= n6; i++ {
		sn += suku
		suku = suku * r
	}

	persen := (sn / sth) * 100

	fmt.Printf("\n*Soal 6 : Analisis Kedekatan Deret Geometri\n")
	fmt.Printf("Input Paket: a = %g, r = %g, N = %d\n", a, r, n6)
	fmt.Printf("Sum Berhingga S(%d): %.2f\n", n6, sn)
	fmt.Printf("Sum Tak Hingga S(inf): %.2f\n", sth)
	fmt.Printf("Persentase Kedekatan: %.2f%%\n", persen)

}
