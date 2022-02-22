package main

import "fmt"

func main() {
	type Weekday int
	const (
		Minggu Weekday = iota
		Senin
		Selasa
		Rabu
		Kamis
		Jumat
		Sabtu
	)
	fmt.Println(Minggu, Senin, Sabtu)
}
