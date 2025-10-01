package assignment_4

import "fmt"

type Truck struct {
}

func (t *Truck) DeliverItem() {
	fmt.Println("Delivering with truck")
}
