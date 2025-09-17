package assignment_2_1

import "fmt"

func ShowPersonDetails() {
	temirlan, _ := getPerson("Temirlan")
	erhan, _ := getPerson("Erhan")

	printDetails(temirlan)
	printDetails(erhan)
}

func printDetails(p IPerson) {
	fmt.Printf("Name: %s", p.getName())
	fmt.Println()
	fmt.Printf("Age: %d", p.getAge())
	fmt.Println()
}
