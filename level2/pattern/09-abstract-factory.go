package main

import (
	"errors"
	"fmt"
)

const (
	BrandAsus = "ASUS"
	BrandHp   = "HP"
)

// computer
type Computer interface {
	PrintInfo()
}

type ComputerAsus struct {
	memory int
	cpu    int
}

func (c *ComputerAsus) PrintInfo() {
	fmt.Println("Info:", c)
}

type ComputerHp struct {
	memory int
	cpu    int
}

func (c *ComputerHp) PrintInfo() {
	fmt.Println("Info:", c)
}

// monitor
type Monitor interface {
	PrintInfo()
}

type MonitorAsus struct {
	size int
}

func (m *MonitorAsus) PrintInfo() {
	fmt.Println("Info:", m)
}

type MonitorHp struct {
	size int
}

func (m *MonitorHp) PrintInfo() {
	fmt.Println("Info:", m)
}

// factory for asus
type AsusFactory struct{}

func (f *AsusFactory) GetComputer() Computer {
	return &ComputerAsus{memory: 8, cpu: 4}
}

func (f *AsusFactory) GetMonitor() Monitor {
	return &MonitorAsus{size: 1500}
}

// factory for hp
type HpFactory struct{}

func (f *HpFactory) GetComputer() Computer {
	return &ComputerHp{memory: 16, cpu: 8}
}

func (f *HpFactory) GetMonitor() Monitor {
	return &MonitorHp{size: 1700}
}

// abstract factory for all
type Factory interface {
	GetComputer() Computer
	GetMonitor() Monitor
}

func GetFactory(brand string) (Factory, error) {
	switch brand {
	default:
		return nil, errors.New(fmt.Sprintf("Brand %s not found!\n", brand))
	case BrandAsus:
		return &AsusFactory{}, nil
	case BrandHp:
		return &HpFactory{}, nil
	}
}

func main() {
	brands := []string{BrandAsus, BrandHp, "DELL"}
	for _, brand := range brands {
		factory, err := GetFactory(brand)
		if err != nil {
			println(err.Error())
			continue
		}
		comp := factory.GetComputer()
		mon := factory.GetMonitor()
		comp.PrintInfo()
		mon.PrintInfo()
	}
}
