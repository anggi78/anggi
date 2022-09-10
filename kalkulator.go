package main

import "fmt"

func tambah(a, b int) int{
	return a + b
}

func kurang(a, b int) int{
	return a - b
}

func kali(a, b int) int{
	return a * b
}

func bagi(a, b float64) float64{
	return a / b
}

func LuasSegitiga(a, t float64) float64{
	return (a * t)/2
}

func LuasPersegi(p, l int) int{
	return p * l
}

func LuasJajarGenjang(a, t int) int{
	return a * t
}

func main() {
	tugasA := tambah(4, 4)
	tugasB := kurang(9, 12)
	tugasC := kali(11, 8)
	tugasD := bagi(5, 6)
	tugasE := LuasSegitiga(10, 7)
	tugasF := LuasPersegi(4, 9)
	tugasG := LuasJajarGenjang(3, 7)

	fmt.Println("Kalkulator")
	fmt.Println(tugasA)
	fmt.Println(tugasB)
	fmt.Println(tugasC)
	fmt.Println(tugasD)
	fmt.Println("Rumus")
	fmt.Println(tugasE)
	fmt.Println(tugasF)
	fmt.Println(tugasG)
}