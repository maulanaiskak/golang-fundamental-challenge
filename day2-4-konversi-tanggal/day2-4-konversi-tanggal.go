package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		tanggal, bulan, tahun int
		bulan_str             string
	)

	fmt.Println("\nKonversi Format Tanggal")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println()

	fmt.Print("Masukkan Tanggal\t: ")
	//errTgl, errBln, errThn untuk memastikan tipe data berupa integer
	_, errTgl := fmt.Scanln(&tanggal) 

	switch {
	case tanggal < 1 || tanggal > 31 || errTgl != nil:
		fmt.Printf("\nTanggal harus bilangan bulat di antara 1 - 31\n\n")

	default:
		fmt.Print("Masukkan Bulan\t\t: ")
		_, errBln := fmt.Scanln(&bulan)

		switch {
		case bulan < 1 || bulan > 12 || errBln != nil:
			fmt.Printf("\nBulan harus bilangan bulat di antara 1 - 12\n\n")

		default:
			// ubah angka bulan ke string
			switch bulan {
			case 1:
				bulan_str = "Januari"

			case 2:
				bulan_str = "Februari"

			case 3:
				bulan_str = "Maret"

			case 4:
				bulan_str = "April"

			case 5:
				bulan_str = "Mei"

			case 6:
				bulan_str = "Juni"

			case 7:
				bulan_str = "Juli"

			case 8:
				bulan_str = "Agustus"

			case 9:
				bulan_str = "September"

			case 10:
				bulan_str = "Oktober"

			case 11:
				bulan_str = "November"

			default:
				bulan_str = "Desember"
			}

			switch {
			case (bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11) && tanggal > 30:
				fmt.Printf("\n%s hanya mempunyai 30 hari\n\n", bulan_str)

			case bulan == 2 && tanggal > 29:
				fmt.Printf("\n%s hanya mempunyai 28/29 hari\n\n", bulan_str)

			default:
				fmt.Print("Masukkan Tahun\t\t: ")
				_, errThn := fmt.Scanln(&tahun)

				switch {
				case tahun < 0 || errThn != nil:
					fmt.Printf("\nTahun harus bilangan bulat positif\n\n")

				default:
					//cek tahun kabisat untuk tanggal bulan februari
					//syarat kabisat : habis dibagi 400 atau tak habis dibagi 400, tak habis dibagi 100 tapi habis dibagi 4//
					bagi400 := tahun%400 == 0 
					bagi100 := tahun%100 == 0
					bagi4 := tahun%4 == 0
					kabisat := bagi400 || (!bagi100 && bagi4)
					
					switch {
					case bulan == 2 && !kabisat:
						fmt.Printf("\n%d bukan tahun kabisat, %s hanya mempunyai 28 hari\n\n", tahun, bulan_str)
					default:
						fmt.Printf("\n%d %s %d\n\n", tanggal, bulan_str, tahun)
					}
				}
			}
		}
	}
}
