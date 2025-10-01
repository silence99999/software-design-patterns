package assignment_4

import "fmt"

type Drone struct {
}

func (d *Drone) DeliverItem() {
	fmt.Println("Delivering with drone")
}
