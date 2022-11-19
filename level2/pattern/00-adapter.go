//Реализовать паттерн «адаптер» на любом примере.

package main

import (
	"fmt"
)

// The "TargetProcess" is the interface that the client wants to call and invoke
type TargetProcess interface {
	Process()
}

// Class that needs to be adapted
type ToBeAdapted struct {
	adapterType   string
	adapterStatus bool
}

// To set adapter status
func (b *ToBeAdapted) SetStatus(status bool) {
	b.adapterStatus = status
}

// To get adapter status
func (b *ToBeAdapted) GetStatus() bool {
	return b.adapterStatus
}

// To get adapter type
func (b *ToBeAdapted) GetType() string {
	return b.adapterType
}

// Adapter struct
type Adapter struct {
	// Object that needs to be adapted
	toBeAdapted ToBeAdapted
}

// Adapter class method process
func (a *Adapter) Process() {
	if a.toBeAdapted.GetStatus() {
		fmt.Println("Adapter", a.toBeAdapted.GetType(), "is already working")
	} else {
		a.toBeAdapted.SetStatus(true)
		fmt.Println("Adapter", a.toBeAdapted.GetType(), "status changed to:", a.toBeAdapted.GetStatus())
	}
}

func main() {
	// Interface object
	var processor TargetProcess = &Adapter{
		toBeAdapted: ToBeAdapted{
			adapterType:   "ADPTR-001",
			adapterStatus: false,
		},
	}

	processor.Process()
}
