package task2

import "fmt"

type Projector struct {
}

func (p *Projector) On() {
	fmt.Println("Projector is ON")
}

func (p *Projector) Off() {
	fmt.Println("Projector is OFF")
}
