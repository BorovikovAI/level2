package main

import "fmt"

const (
	ComputerBrandAsus = "asus"
	ComputerBrandHp   = "hp"
)

type Computer struct {
	core   int
	memory int
	ram    int
}

func (c *Computer) Print() {
	fmt.Println(c)
}

type ComputerBuilder interface {
	SetCore()
	SetMemory()
	SetRam()
	BuildComputer() *Computer
}

func NewComputerBuider(brand string) *ComputerBuilder {
	switch brand {
	default:
		fmt.Printf("%s no such brand!", brand)
		return nil
	case ComputerBrandAsus:
		return BuildAsus()
	case ComputerBrandHp:
		return BuildHp()
	}
}

type ComputerAsus struct {
	core   int
	memory int
	ram    int
}

func BuildAsus() *ComputerBuilder {
	var c ComputerBuilder = &ComputerAsus{}
	return &c
}

func (c *ComputerAsus) SetCore() {
	c.core = 4
}

func (c *ComputerAsus) SetMemory() {
	c.memory = 256
}

func (c *ComputerAsus) SetRam() {
	c.ram = 8
}

func (c *ComputerAsus) BuildComputer() *Computer {
	return &Computer{
		core:   c.core,
		memory: c.memory,
		ram:    c.ram,
	}
}

type ComputerHp struct {
	core   int
	memory int
	ram    int
}

func BuildHp() *ComputerBuilder {
	var c ComputerBuilder = &ComputerHp{}
	return &c
}

func (c *ComputerHp) SetCore() {
	c.core = 8
}

func (c *ComputerHp) SetMemory() {
	c.memory = 512
}

func (c *ComputerHp) SetRam() {
	c.ram = 16
}

func (c *ComputerHp) BuildComputer() *Computer {
	return &Computer{
		core:   c.core,
		memory: c.memory,
		ram:    c.ram,
	}
}

type ComputerFactory struct {
	computerBuilder ComputerBuilder
}

func NewFactory(computerBuilder ComputerBuilder) *ComputerFactory {
	return &ComputerFactory{computerBuilder: computerBuilder}
}

func (f *ComputerFactory) SetComputerBuilder(computerBuilder ComputerBuilder) {
	f.computerBuilder = computerBuilder
}

func (f *ComputerFactory) CreateComputer() *Computer {
	f.computerBuilder.SetCore()
	f.computerBuilder.SetMemory()
	f.computerBuilder.SetRam()
	return f.computerBuilder.BuildComputer()
}

func main() {
	asus := NewComputerBuider(ComputerBrandAsus)
	hp := NewComputerBuider(ComputerBrandHp)

	factory := NewFactory(*asus)
	computer := factory.CreateComputer()
	computer.Print()

	factory.SetComputerBuilder(*hp)
	computer = factory.CreateComputer()
	computer.Print()
}
