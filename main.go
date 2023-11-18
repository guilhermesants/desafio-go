package main

import "fmt"

func main() {

	var temperatura float64
	fmt.Print("Informe a temperatura: ")
	fmt.Scan(&temperatura)

	fmt.Printf("%.2f Kelvin é igual a %.2f Celsius.", temperatura, kelvinParaCelsius(temperatura))
}

func kelvinParaCelsius(kelvinValue float64) float64 {
	return kelvinValue - 273.00
}
