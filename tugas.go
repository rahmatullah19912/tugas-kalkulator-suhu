package main

import (
	"fmt"
)

func main() {
	var a string
	var konv int32
	var suhu_awal, suhu_akhir float32
	a = "y"

	fmt.Println("__________________________________________")
	fmt.Println("Selamat menggunakan aplikasi konversi suhu")
	fmt.Println("__________________________________________")
	fmt.Println(" ")

	for a != "n" {
		fmt.Print("Apakah anda akan melakukan konversi suhu ( y / n) = ")
		fmt.Scan(&a)

		if a == "y" {
			fmt.Println("")
			fmt.Println("________________________________________________")
			fmt.Println("Pilihlah jenis konversi dibawah ini : ")
			fmt.Println("")
			fmt.Println("1. Celcius to Reamur")
			fmt.Println("2. Celcius to Fahrenheit")
			fmt.Println("3. celcius to kelvin")
			fmt.Println("4. Kelvin to Celcius")
			fmt.Println("5. Kelvin to Reamur")
			fmt.Println("6. Kelvin to Fahreinheit")
			fmt.Println("7. Reamur to Celcius")
			fmt.Println("8. Reamur to Kelvin")
			fmt.Println("9. Reamur to Fahrenheit")
			fmt.Println("10.Fahrenheit to Celcius")
			fmt.Println("11.Fahrenheit to Kelvin")
			fmt.Println("12.Fahrenheit to Reamur")
			fmt.Println("_______________________________________")
			fmt.Println("")
			fmt.Print("input jenis konversi : ")
			fmt.Scan(&konv)

			fmt.Print("input suhu yang akan dikonversi : ")
			fmt.Scan(&suhu_awal)

			if konv == 1 {
				suhu_akhir = 0.8 * suhu_awal
			} else if konv == 2 {
				suhu_akhir = 1.8*suhu_awal + 32
			} else if konv == 3 {
				suhu_akhir = 273 + suhu_awal
			} else if konv == 4 {
				suhu_akhir = suhu_awal - 273
			} else if konv == 5 {
				suhu_akhir = 0.8 * (suhu_awal - 273)
			} else if konv == 6 {
				suhu_akhir = 1.8*(suhu_awal-273) + 32
			} else if konv == 7 {
				suhu_akhir = 1.25 * suhu_awal
			} else if konv == 8 {
				suhu_akhir = 1.23*suhu_awal + 273
			} else if konv == 9 {
				suhu_akhir = 2.25*suhu_awal + 32
			} else if konv == 10 {
				suhu_akhir = 0.5556 * (suhu_awal + 32)
			} else if konv == 11 {
				suhu_akhir = 0.5556*(suhu_awal-32) + 273
			} else if konv == 12 {
				suhu_akhir = 0.444 * (suhu_awal - 32)
			} else {
				fmt.Print("Jenis konversi tidak terdefinisi, mulai dari awal lagi ya. ")
			}

			fmt.Println("Hasil ")
			fmt.Print(suhu_awal)

			fmt.Print(" = ")
			fmt.Println(suhu_akhir)
			fmt.Println("__________________________________________")

		} else if a != "n" {
			fmt.Println("Input hanya -y- untuk YES atau -n- untuk NO")

		}
	}
	fmt.Println("Terimakasih :)")

}
