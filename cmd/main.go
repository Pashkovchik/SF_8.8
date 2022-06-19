package main

import (
	electronic "SF_8.8/pkg"
	"fmt"
)

func printCharacteristics(p electronic.Phone) {
	fmt.Println("Бренд телефона:", p.Brand())
	fmt.Println("Модель телефона:", p.Model())
	fmt.Println("Тип телефона:", p.Type())
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
