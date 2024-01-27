package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type institute struct {
	instituteName          string
	instituteAddress       string
	instituteAccreditation string
	instituteSince         int
}

type course struct {
	name string
}

type student struct {
	npm, name string
	ins       institute
	courses   []course
}

func main() {
	fmt.Println("Aplikasi Data Mahasiswa")
	fmt.Println(strings.Repeat("-", 80))

	var npm, name, instituteName, instituteAddress, instituteAccreditation, courses, operation, showBy, again string
	var instituteSince, idx int
	var coursesSplit []string
	var coursesSlice []course
	var dataSlice []student

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nOperation (add/delete/show/exit)?: ")
		scanner.Scan()
		operation = scanner.Text()

		if operation == "add" {
			if len(dataSlice) <= 5 {
				fmt.Print("\nNPM\t\t\t\t: ")
				scanner.Scan()
				npm = scanner.Text()

				i := false
				for !i {
					fmt.Print("Nama\t\t\t\t: ")
					scanner.Scan()
					name = scanner.Text()

					if len(name) < 3 || len(name) > 20 {
						i = false
						fmt.Printf("\nNama Harus 3-20 Karakter, Coba Lagi\n\n")
					} else {
						i = true
					}
				}

				fmt.Print("Nama Institusi\t\t\t: ")
				scanner.Scan()
				instituteName = scanner.Text()

				fmt.Print("Alamat Institusi\t\t: ")
				scanner.Scan()
				instituteAddress = scanner.Text()

				fmt.Print("Akreditasi Institusi\t\t: ")
				scanner.Scan()
				instituteAccreditation = scanner.Text()

				i = false
				for !i {
					fmt.Print("Tahun Pendirian Institusi\t: ")
					_, err := fmt.Scanln(&instituteSince)

					if err != nil {
						i = false
						fmt.Printf("\nTahun Salah, Coba Lagi\n\n")
					} else {
						i = true
					}
				}

				fmt.Print("Mata Kuliah (MK1,MK2,MK3...)\t: ")
				scanner.Scan()
				courses = scanner.Text()
				coursesSplit = strings.Split(courses, ",")

				coursesSlice = []course{}
				for _, value := range coursesSplit {
					coursesSlice = append(coursesSlice, course{value})
				}

				fmt.Printf("\nData Berhasil Ditambahkan\n")
			} else {
				fmt.Println("Database Hanya bisa Menampung 5 Data, Hapus Data Lain Terlebih Dahulu")
				continue
			}
		} else if operation == "delete" {
			if len(dataSlice) == 0 {
				fmt.Println("\nData Kosong, Tambah Data Terlebih Dahulu")
				continue
			} else {
				idx = validIdx(len(dataSlice), idx)
			}
			fmt.Printf("\nData dengan Index %d Berhasil Dihapus\n", idx)
		} else if operation == "show" {
			if len(dataSlice) == 0 {
				fmt.Println("\nData Kosong, Tambah Data Terlebih Dahulu")
				continue
			} else {
				showBy = validOption("\nShow all (all) / by index(idx)) ? : ", "all", "idx")
				if showBy == "idx" {
					idx = validIdx(len(dataSlice), idx)
				}
			}
		} else if operation == "exit" {
			break

		} else {
			fmt.Println("\nInput Salah, Coba Lagi")
			continue
		}

		dataSlice = database(dataSlice, operation, showBy, npm, name, instituteName, instituteAddress, instituteAccreditation, instituteSince, idx, coursesSlice)

		again = validOption("\nLakukan Operasi Database Lagi (yes/no)?", "yes", "no")
		if again == "yes" {
			continue
		} else {
			break
		}
	}
}

func database(dataSlice []student, operation, showBy, npm, name, instituteName, instituteAddress, instituteAccreditation string, insSince, idx int, courses []course) []student {
	if operation == "add" {
		data := student{
			npm,
			name,
			institute{
				instituteName,
				instituteAddress,
				instituteAccreditation,
				insSince,
			},
			courses,
		}

		dataSlice = append(dataSlice, data)
	} else if operation == "delete" {
		dataSlice = append(dataSlice[:idx], dataSlice[idx+1:]...)
	} else if operation == "show" {
		if showBy == "idx" {
			fmt.Printf("Index \t\t: %d\n", idx)
			fmt.Printf("NPM \t\t: %s\n", dataSlice[idx].npm)
			fmt.Printf("Nama \t\t: %s\n", dataSlice[idx].name)
			fmt.Printf("Institusi \t: %s, %s, %s, %d\n", dataSlice[idx].ins.instituteName, dataSlice[idx].ins.instituteAddress, dataSlice[idx].ins.instituteAccreditation, dataSlice[idx].ins.instituteSince)
			fmt.Printf("Mata Kuliah \t: ")
			for _, j := range dataSlice[idx].courses {
				fmt.Printf("%s, ", j.name)
			}
			fmt.Println()
		} else {
			for id, i := range dataSlice {
				fmt.Printf("Index \t\t: %d\n", id)
				fmt.Printf("NPM \t\t: %s\n", i.npm)
				fmt.Printf("Nama \t\t: %s\n", i.name)
				fmt.Printf("Institusi \t: %s, %s, %s, %d\n", i.ins.instituteName, i.ins.instituteAddress, i.ins.instituteAccreditation, i.ins.instituteSince)
				fmt.Printf("Mata Kuliah \t: ")
				for _, j := range i.courses {
					fmt.Printf("%s, ", j.name)
				}
				fmt.Printf("\n\n")
			}
		}
	}
	return dataSlice
}

func validOption(question, optionA, optionB string) string {
	var answer string
	i := false
	for !i {
		fmt.Print(question)
		fmt.Scanln(&answer)

		if answer == optionA || answer == optionB {
			i = true
		} else {
			fmt.Println("\nInput Salah, Coba Lagi")
			i = false
		}
	}

	return answer
}

func validIdx(length, idx int) int {
	i := false
	for !i {
		fmt.Print("\nMasukkan index: ")
		fmt.Scanln(&idx)

		if idx < length && idx > 0 {
			i = true
		} else {
			fmt.Printf("\nIndex Salah, Masukkan 0 - %d\n", length-1)
			i = false
		}
	}
	return idx
}
