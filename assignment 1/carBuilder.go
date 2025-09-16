package main

import (
	"fmt"
)

type BuildPC struct {
	cpu     string
	gpu     string
	coolers int
	ram     string
}

type Builder struct {
	cpu     string
	gpu     string
	coolers int
	ram     string
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) SetCPUModel(cpu string) *Builder {
	b.cpu = cpu
	return b
}

func (b *Builder) SetGPUModel(gpu string) *Builder {
	b.gpu = gpu
	return b
}

func (b *Builder) SetCoolersAmount(coolers int) *Builder {
	b.coolers = coolers
	return b
}

func (b *Builder) SetRAMAmount(ram string) *Builder {
	b.ram = ram
	return b
}

func (b *Builder) Build() *BuildPC {
	return &BuildPC{
		cpu:     b.cpu,
		gpu:     b.gpu,
		coolers: b.coolers,
		ram:     b.ram,
	}
}

func (pc *BuildPC) String() string {
	return fmt.Sprintf("BuildPC { CPU='%s', GPU='%s', Coolers=%d, RAM='%s' }", pc.cpu, pc.gpu, pc.coolers, pc.ram)
}
