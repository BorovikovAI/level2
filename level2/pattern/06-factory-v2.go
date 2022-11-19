package main

import (
	"fmt"
)

const (
	ComputerTypeServer   = "server"
	ComputerTypePersonal = "personal"
)

// computer
type Computer interface {
	GetType()
	GetAll()
}

func SetComputer(name string) Computer {
	switch name {
	default:
		fmt.Printf("%s no such type!\n", name)
		return nil
	case ComputerTypeServer:
		fmt.Printf("New Comp: %s\n", name)
		return NewServer()
	case ComputerTypePersonal:
		fmt.Printf("New Comp: %s\n", name)
		return NewPersonal()
	}
}

// server computer
type Server struct {
	code   string
	core   int
	memory int
}

func NewServer() Computer {
	return &Server{
		code:   "server_model",
		core:   16,
		memory: 256,
	}
}

func (s *Server) GetType() {
	fmt.Println(s.code)
}

func (s *Server) GetAll() {
	fmt.Println(s)
}

// personal computer
type Personal struct {
	code   string
	core   int
	memory int
}

func NewPersonal() Computer {
	return &Personal{
		code:   "personal_model",
		core:   8,
		memory: 128,
	}
}

func (p *Personal) GetType() {
	fmt.Println(p.code)
}

func (p *Personal) GetAll() {
	fmt.Println(p)
}

// factory
type Factory struct {
	computer Computer
}

func NewFactory(computer Computer) *Factory {
	fmt.Print("New Factory: ")
	computer.GetType()
	return &Factory{computer: computer}
}

func (f *Factory) SetFactory(compType string) *Factory {
	newComputer := SetComputer(compType)
	if newComputer == nil {
		return nil
	}
	newComputer.GetAll()
	return NewFactory(newComputer)
}

// main
func main() {
	// SetComputer(ComputerTypeServer)
	// SetComputer(ComputerTypePersonal)
	// SetComputer("monoblock")
	comp := SetComputer(ComputerTypeServer)
	factory := NewFactory(comp)
	factory.SetFactory(ComputerTypePersonal)
	factory.SetFactory("monoblock")
}
