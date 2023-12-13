package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//Menerima Input
	var text, temp2 string
	var temp, n, m, errNo int
	var errCheck error
	total := 0
	errorDetect := ""
	fmt.Print("Silakan Input String Bilangan: ")
	fmt.Scanln(&text)

	//Perhitungan dan pengecekan input
	if strings.Contains(string(text[0]), "+") || strings.Contains(string(text[0]), "-") {
		for i := 0; i < len(text); i++ {
			//Pengecekan string memiliki bilangan
			for x := 0; x < 10; x++ {
				if strings.Contains(string(text[i]), fmt.Sprint(x)) {
					
					if i+1 == len(text) {
						//Cek sudah tidak ada lagi karakter setelah ini
						temp, errCheck = strconv.Atoi(string(text[i]))
						//Error di Value
						if errCheck != nil {
							//fmt.Println(errNo) //Pengecekan Jumlah Error
							errCheck = nil
						}
					} else if strings.Contains(string(text[i-1]), "+") || strings.Contains(string(text[i-1]), "-"){
						
						//Jika karakter sebelumnya bilangan atau karakter lain, tidak lanjut
						//(Agar setiap angka hanya masuk 1x)
						for n = i; n < len(text); n++{
							if n+1 == len(text) || string(text[n+1]) == "+" || string(text[n+1]) == "-" {
								if n-i == 0 && n+1 != len(text) {
									//Cek apakah 1 digit saja
									temp, errCheck = strconv.Atoi(string(text[i]))
									//Error di Value
									if errCheck != nil {
										//fmt.Println(errNo) //Pengecekan Jumlah Error
										errCheck = nil
									}
								} else {
									//Lebih dari 1 digit
									temp2 = ""
									for m = 0; m <= n-i; m++ {
										temp2 += string(text[i+m])
										//fmt.Println(temp2,"|",m) //Pengecekan nilai temp
									}
									temp, errCheck = strconv.Atoi(string(temp2))
									//Error di Value
									if errCheck != nil {
										//fmt.Println(errNo) //Pengecekan Jumlah Error
										errCheck = nil
									}
								}
								//Sudah ketemu + atau - berikut
								break
							} 
						}					
					}
					
					if string(text[i-1]) == "+" {
						//Bilangan Positif					
						total = total + temp
						//fmt.Println(total) //Pengecekan Total Progress
					} else if string(text[i-1]) == "-" {
						//Bilangan Negatif
						total = total - temp
						//fmt.Println(total) //Pengecekan Total Progress
					}
				} 
			}
			if string(text[i]) == "+" || string(text[i]) == "-" {
				
			} else {
				_, errCheck = strconv.Atoi(string(text[i]))
				//Error di Value
				if errCheck != nil {
					errNo++
					//fmt.Println(errNo,"a") //Pengecekan Jumlah Error
					errCheck = nil
				}
			}			
		}
		//Output
		if errNo != 0 {
			//Error Handling
			errorDetect = " [WARNING] " + strconv.Itoa(errNo) + " Error terdeteksi, mohon tidak menginput huruf atau karakter selain bilangan dan operator + dan -"
		}
		fmt.Print(total, errorDetect)
	} else {
		fmt.Print("Mohon mengisi baris dengan deret +(Bilangan) atau -(Bilangan)")
	}
}