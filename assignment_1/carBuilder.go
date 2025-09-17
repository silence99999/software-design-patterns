package assignment_1

import (
	"fmt"
)

type PC struct {
	cpu     string
	gpu     string
	coolers int
	ram     string
}

type PCBuilder struct {
	pc PC
}

func NewBuilder() *PCBuilder {
	return &PCBuilder{}
}

func (b *PCBuilder) SetCPUModel(cpu string) *PCBuilder {
	b.pc.cpu = cpu
	return b
}

func (b *PCBuilder) SetGPUModel(gpu string) *PCBuilder {
	b.pc.gpu = gpu
	return b
}

func (b *PCBuilder) SetCoolersAmount(coolers int) *PCBuilder {
	b.pc.coolers = coolers
	return b
}

func (b *PCBuilder) SetRAMAmount(ram string) *PCBuilder {
	b.pc.ram = ram
	return b
}

func (b *PCBuilder) Build() *PC {
	return &b.pc
}

func (pc *PC) String() string {
	return fmt.Sprintf("BuildPC { CPU='%s', GPU='%s', Coolers=%d, RAM='%s' }", pc.cpu, pc.gpu, pc.coolers, pc.ram)
}
