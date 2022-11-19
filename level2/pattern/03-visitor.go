package main

import (
	"fmt"
)

type Cocktail interface {
	accept(v CocktailVisitor)
}

type CocktailVisitor interface {
	visitMojito(c Mojito)
	visitDaiquiri(c Daiquiri)
}

type DrinkVisit struct{}

func (DrinkVisit) visitMojito(m Mojito) {
	fmt.Println("Drinking mojitos")
}
func (DrinkVisit) visitDaiquiri(d Daiquiri) {
	fmt.Println("Drinking daiquiris")
}

type ServeVisit struct{}

func (ServeVisit) visitMojito(m Mojito) {
	fmt.Println("Serving mojitos")
}
func (ServeVisit) visitDaiquiri(d Daiquiri) {
	fmt.Println("Serving daiquiris")
}

type Mojito struct{}

func (m Mojito) accept(v CocktailVisitor) {
	v.visitMojito(m)
}

type Daiquiri struct{}

func (d Daiquiri) accept(v CocktailVisitor) {
	v.visitDaiquiri(d)
}

func main() {
	// creating Cocktail interface objects
	mojito := &Mojito{}
	daiquiri := &Daiquiri{}
	// creating CocktailVisitor interface objects
	drink := &DrinkVisit{}
	serve := &ServeVisit{}

	mojito.accept(drink)
	daiquiri.accept(serve)
}

// The visitor pattern is a common pattern in language interpreters
// that allows the interpreter to work on various expressions
// by defining operations on them. Typically, the accept method
// will return an interface{} or a value.
