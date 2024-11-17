package main

import "fmt"

func main() {
	var kelvin float64

	fmt.Print("Digite o valor da temperatura em Kelvin: ")
	fmt.Scan(&kelvin)

	celsius := kelvin - 273.15

	fmt.Printf("O valor da temperatura em Celsius e: %.2f°C\n", celsius)
}
