package main

import "fmt"

func main() {
	var input float64
	fmt.Print("Input temperature in Fahrenheit: ")
	fmt.Scanf("%f", &input)
	fahrenheitToCelsiusConverter(input)
}

func fahrenheitToCelsiusConverter(temp float64) {
	celsius := (temp - 32) * 5 / 9
	// strCelcius := strconv.FormatFloat(celsius, 'f', 3, 64)
	// strFahrenheit := strconv.FormatFloat(temp, 'f', 3, 64)
	fmt.Printf("%.2f", celsius)
	fmt.Print("°F is ")
	fmt.Printf("%.2f", temp)
	fmt.Println("°C")
	// fmt.Println(strFahrenheit + "°F is " + strCelcius + "°C")
}
