package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			saida := fmt.Sprintf(`o numero %v é divisivel por 3`, i)
			fmt.Println(saida)
		}
	}
}
