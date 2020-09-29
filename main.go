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
			fmt.Print("acuario")
		} else {
			fmt.Print("pisis")
		}
	case 3:
		if dia < 21 {
			fmt.Print("pisis")
		} else {
			fmt.Print("aries")
		}
	case 4:
		if dia < 20 {
			fmt.Print("aries")
		} else {
			fmt.Print("tauro")
		}
	case 5:
		if dia < 21 {
			fmt.Print("tauro")
		} else {
			fmt.Print("geminis")
		}
	case 6:
		if dia < 21 {
			fmt.Print("geminis")
		} else {
			fmt.Print("cancer")
		}
	case 7:
		if dia < 23 {
			fmt.Print("cancer")
		} else {
			fmt.Print("leo")
		}
	case 8:
		if dia < 23 {
			fmt.Print("leo")
		} else {
			fmt.Print("virgo")
		}
	case 9:
		if dia < 23 {
			fmt.Print("virgo")
		} else {
			fmt.Print("libra")
		}
	case 10:
		if dia < 23 {
			fmt.Print("libra")
		} else {
			fmt.Print("escorpio")
		}
	case 11:
		if dia < 22 {
			fmt.Print("escorpio")
		} else {
			fmt.Print("sagitario")
		}
	case 12:
		if dia < 22 {
			fmt.Print("sagitario")
		} else {
			fmt.Print("capricornio")
		}
	default:
		fmt.Print("error")
	}
}
