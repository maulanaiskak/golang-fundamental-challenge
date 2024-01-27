package main

import (
	"fmt"
	"strings"
)

func main() {
	var jam int

	fmt.Println("\nJadwal Harian Susi")
	fmt.Printf("%s\n\n", strings.Repeat("-", 50))

	fmt.Printf("Masukkan Jam: ")
	_, err := fmt.Scanln(&jam)
	fmt.Println()

	//memastikan jam di antara 0-24, jam negatif -> eror, jam > 24 -> eror
	if jam >= 0 && jam <= 24 && err == nil {
		if jam >= 4 && jam <= 5 {
			fmt.Printf("Jam %d Susi Bangun Pagi\n", jam)
		} else if jam >= 6 && jam <= 7 {
			fmt.Printf("Jam %d Susi Mandi dan Sarapan\n", jam)
		} else if jam >= 8 && jam <= 11 {
			fmt.Printf("Jam %d Susi Berangkat Sekolah\n", jam)
		} else if jam == 12 {
			fmt.Printf("Jam %d Susi Pulang Sekolah\n", jam)
		} else if jam >= 13 && jam <= 15 {
			fmt.Printf("Jam %d Susi Tidur Siang\n", jam)
		} else if jam >= 16 && jam <= 17 {
			fmt.Printf("Jam %d Susi Bermain\n", jam)
		} else {
			fmt.Printf("Jam %d Susi Istirahat\n", jam)
		}
	} else {
		fmt.Println("Jam Salah, Hanya Ada 24 Jam dalam Sehari")
	}
	fmt.Println()

}
