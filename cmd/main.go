package main

import (
	electronic "SF_8.8/pkg"
	"fmt"
)

func printCharacteristics(p electronic.Phone) {
	pis, _ := p.(electronic.Smartphone)
	pist, _ := p.(electronic.StationPhone)
	fmt.Printf("Brand: %q Model: %q Type Of Phone: %q ", p.Brand(), p.Model(), p.Type())
	switch p.Type() {
	case "smartphone":
		fmt.Printf("OS: %q\n", pis.OS())
	case "station":
		fmt.Printf("Buttons Count: %d\n", pist.ButtonsCount())
	default:
		panic("unknown type of phone")
	}
}

func main() {
	myIphone := electronic.NewApplePhone("12 Pro")
	fmt.Println(myIphone)
	mySamsung := electronic.NewAndroidPhone("Galaxy", "SAMSUNG")
	myStat := electronic.NewRadioPhone("S500", "Витязь", 3)
	fmt.Println("Вся информация об Iphone:")
	printCharacteristics(myIphone)
	fmt.Println("")
	fmt.Println("Вся информация о Samsung:")
	printCharacteristics(mySamsung)
	fmt.Println("")
	fmt.Println("Вся информация о Витязе:")
	printCharacteristics(myStat)
}
