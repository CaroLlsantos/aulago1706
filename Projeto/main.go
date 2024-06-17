package main

import "fmt"

func main() {

	menu := map[string]float64{
		"Pizza": 40.00,
		"Suco": 12.50,
		"X-tudo": 22.90,
	}

	for k,v := range menu{
		fmt.Println(k, " R$", v)
	}
	contanova := novaConta("Astrubal")
	fmt.Println(contanova)

	contanova2 := novaConta("Jubiel")
	fmt.Println(contanova2)	
}