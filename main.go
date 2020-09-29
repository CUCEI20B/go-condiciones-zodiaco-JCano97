package main

import "fmt"

func main() {
	var dia uint64
	var mes uint64

	fmt.Scan(&dia)
	fmt.Scan(&mes)

	switch mes {
	case 1:
		if dia < 20 {
			fmt.Print("capricornio")
		} else {
			fmt.Print("acuario")
		}
	case 2:
		if dia < 19 {
			fmt.Print("Acuario")
		} else {
			fmt.Print("Pisis")
		}
	case 3:
		if dia < 21 {
			fmt.Print("Pisis")
		} else {
			fmt.Print("Aries")
		}
	case 4:
		if dia < 20 {
			fmt.Print("Aries")
		} else {
			fmt.Print("Tauro")
		}
	case 5:
		if dia < 21 {
			fmt.Print("Tauro")
		} else {
			fmt.Print("Géminis")
		}
	case 6:
		if dia < 21 {
			fmt.Print("Géminis")
		} else {
			fmt.Print("Cáncer")
		}
	case 7:
		if dia < 23 {
			fmt.Print("Cáncer")
		} else {
			fmt.Print("Leo")
		}
	case 8:
		if dia < 23 {
			fmt.Print("Leo")
		} else {
			fmt.Print("Virgo")
		}
	case 9:
		if dia < 23 {
			fmt.Print("Virgo")
		} else {
			fmt.Print("Libra")
		}
	case 10:
		if dia < 23 {
			fmt.Print("Libra")
		} else {
			fmt.Print("Escorpio")
		}
	case 11:
		if dia < 22 {
			fmt.Print("Escorpio")
		} else {
			fmt.Print("Sagitario")
		}
	case 12:
		if dia < 22 {
			fmt.Print("Sagitario")
		} else {
			fmt.Print("Capricornio")
		}
	default:
		fmt.Print("error")
	}
}
