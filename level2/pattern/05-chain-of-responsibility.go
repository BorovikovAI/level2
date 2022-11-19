package main

import (
	"fmt"
)

// Handler interface.
// This interface handling requests and sets the next handler on the chain.
type section interface {
	execute(*task)
	setNext(section)
}

type task struct {
	name              string
	materialCollected bool
	assemblyExecuted  bool
	packagingExecuted bool
}

// Concrete handler.
// It decides if the request to collect material should be processed and can move up the chain.
type material struct {
	name string
	next section
}

func (m *material) execute(task *task) {
	fmt.Println(task)
	if task.materialCollected {
		fmt.Println("Material already collected")
		m.next.execute(task)
		return
	}
	fmt.Println("Material section gathering materials")
	task.materialCollected = true
	m.next.execute(task)
}

func (m *material) setNext(next section) {
	m.next = next
}

// Concrete handler.
// It decides if the request to perform the assembly work should be processed and can move up the chain.
type assembly struct {
	name string
	next section
}

func (a *assembly) execute(t *task) {
	fmt.Println(t)
	if t.assemblyExecuted {
		fmt.Println("Assembly already done")
		a.next.execute(t)
		return
	}
	fmt.Println("Assembly section assembling...")
	t.assemblyExecuted = true
	a.next.execute(t)
}

func (a *assembly) setNext(next section) {
	a.next = next
}

// Concrete handler.
// It decides if the request to perform the packaging work should be processed and can move up the chain.
type packaging struct {
	name string
	next section
}

func (p *packaging) execute(t *task) {
	fmt.Println(t)
	if t.packagingExecuted {
		fmt.Println("Packaging already done")
		p.next.execute(t)
		return
	}
	fmt.Println("Packaging section doing packaging")
}

func (p *packaging) setNext(next section) {
	p.next = next
}

// main is a client. Initializes up the chain of handlers.
func main() {
	packaging := &packaging{name: "packaging"}

	// set next for assembly section
	assembly := &assembly{name: "assembly"}
	assembly.setNext(packaging)

	material := &material{name: "material"}
	material.setNext(assembly)

	task := &task{name: "B.R.U.H."}
	material.execute(task)
}
