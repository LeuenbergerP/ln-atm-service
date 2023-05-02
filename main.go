package main

import (
	"fmt"
)

func main() {
	client := NewOpenNodeClient("API_KEY", "dev")
	cha, err := client.CreateCharge(Charge{})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cha)
}
