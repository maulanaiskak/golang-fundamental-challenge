package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type PenjualanToko struct {
	KodeToko  string
	Penjualan []float64
}

func main() {
	var wg sync.WaitGroup
	var mtx sync.Mutex
	var total float64

	timeStart := time.Now()
	dataToko := []PenjualanToko{
		{"123", []float64{1, 2, 3, 4, 5}},
		{"124", []float64{13, 21, 33}},
		{"125", []float64{4, 2}},
		{"126", []float64{21, 22, 23, 9, 1, 22}},
		{"127", []float64{8, 4, 10}},
	}

	fmt.Println("PERHITUNGAN PENJUALAN TOKO")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println()

	for _, toko := range dataToko {
		wg.Add(1)
		go HitungPenjualan(toko, &total, &wg, &mtx)
	}
	wg.Wait()

	timeElapsed := time.Since(timeStart)
	fmt.Printf("\nTotal Keseluruhan Penjualan\t: %.f\n", total)
	fmt.Printf("\nExecution time: %s", timeElapsed)
}

func HitungPenjualan(n PenjualanToko, total *float64, wg *sync.WaitGroup, mtx *sync.Mutex) {
	var totalToko float64
	for _, penjualan := range n.Penjualan {
		mtx.Lock()

		totalToko += penjualan //total per toko
		*total += penjualan // total keseluruhan toko, mutex bekerja untuk variabel ini karena diakses 5 goroutine

		mtx.Unlock()
	}
	wg.Done()

	fmt.Printf("Total penjualan toko %s\t: %.f\n", n.KodeToko, totalToko)
}
