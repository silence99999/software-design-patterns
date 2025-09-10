package main

import (
	"fmt"
)

type BuildPC struct {
	CPU     string
	GPU     string
	Coolers int
	RAM     string
}

type Builder struct {
	cpu     string
	gpu     string
	coolers int
	ram     string
}

func (b *Builder) SetCPU(cpu string) *Builder {
	b.cpu = cpu
	return b
}

func (b *Builder) SetGPU(gpu string) *Builder {
	b.gpu = gpu
	return b
}

func (b *Builder) SetCoolers(coolers int) *Builder {
	b.coolers = coolers
	return b
}

func (b *Builder) SetRAM(ram string) *Builder {
	b.ram = ram
	return b
}

func (b *Builder) Build() *BuildPC {
	return &BuildPC{
		CPU:     b.cpu,
		GPU:     b.gpu,
		Coolers: b.coolers,
		RAM:     b.ram,
	}
}

func (pc *BuildPC) String() string {
	return fmt.Sprintf("BuildPC { CPU='%s', GPU='%s', Coolers=%d, RAM='%s' }",
		pc.CPU, pc.GPU, pc.Coolers, pc.RAM)
}

func main() {
	b := &Builder{}
	pc := b.
		SetCPU("RYZEN 7 7435hs").
		SetGPU("NVIDIA RTX 4050").
		SetCoolers(3).
		SetRAM("24GB DDR6").
		Build()

	fmt.Println(pc)
}
