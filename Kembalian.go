package main

import (
	"fmt"
	"strconv"
)

func main() {
	tes, hasil := Kembalian(18000, 5000)
	fmt.Println(tes, hasil)
}

func Kembalian(uang, bayar int) (string, []string) {
	rupiah := []int{100000, 50000, 25000, 10000, 5000, 1000}
	banyak := 0
	pecahan := 0
	Kembalian := uang - bayar
	hasil := []string{}

	if uang % 1000 != 0 || uang < bayar{
		return "invalid", nil
	}

	for _, i := range rupiah {
		if i > Kembalian {
			continue
		}
		banyak = Kembalian / i
		pecahan = i
		Kembalian = Kembalian - i
		hasil = append(hasil, strconv.Itoa(pecahan) + " ada " + strconv.Itoa(banyak) + ", ")
	}
	return "", hasil
}
