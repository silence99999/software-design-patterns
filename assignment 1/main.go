package main

import "fmt"

func main() {
	pcRyzen7Rtx4050 := NewBuilder().
		SetCPUModel("RYZEN 7 7435hs").
		SetGPUModel("NVIDIA RTX 4050").
		SetCoolersAmount(3).
		SetRAMAmount("24GB DDR6").
		Build()

	fmt.Println(pcRyzen7Rtx4050)
}
