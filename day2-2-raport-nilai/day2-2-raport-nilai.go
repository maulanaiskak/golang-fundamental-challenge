package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		tugas, uts, uas float32
	)

	fmt.Println("\nRaport Nilai Akhir")
	fmt.Printf("%s\n\n", strings.Repeat("-", 50))

	// errTugas, errUTS, errUAS untuk memastikan input berupa bilangan
	// jika terjadi eror -> program berhenti
	fmt.Printf("Masukkan nilai tugas\t: ")
	_, errTugas := fmt.Scanln(&tugas)

	if tugas < 0 || tugas > 100 || errTugas != nil {
		fmt.Println("Nilai tugas harus berupa bilangan desimal 0.0 - 100.0")
	} else {
		fmt.Printf("Masukkan nilai UTS\t: ")
		_, errUTS := fmt.Scanln(&uts)

		if uts < 0 || uts > 100 || errUTS != nil {
			fmt.Println("Nilai UTS harus berupa bilangan desimal 0.0 - 100.0")
		} else {
			fmt.Printf("Masukkan nilai UAS\t: ")
			_, errUAS := fmt.Scanln(&uas)
			fmt.Println()

			if uas < 0 || uas > 100 || errUAS != nil {
				fmt.Println("Nilai UAS harus berupa bilangan desimal 0.0 - 100.0")
			} else {
				if nilai := 0.25*tugas + 0.3*uts + 0.45*uas; nilai >= 80 && nilai <= 100 {
					fmt.Printf("Nilai %.2f dengan hasil akhir adalah A\n", nilai)
				} else if nilai >= 70 && nilai < 80 {
					fmt.Printf("Nilai %.2f dengan hasil akhir adalah B\n", nilai)
				} else if nilai >= 55 && nilai < 70 {
					fmt.Printf("Nilai %.2f dengan hasil akhir adalah C\n", nilai)
				} else if nilai >= 40 && nilai < 55 {
					fmt.Printf("Nilai %.2f dengan hasil akhir adalah D\n", nilai)
				} else {
					fmt.Printf("Nilai %.2f dengan hasil akhir adalah E\n", nilai)
				}
			}
		}
	}
	fmt.Println()

}
