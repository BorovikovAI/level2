package main

import (
	"fmt"
)

// Person struct
type Person struct {
	// Personal details
	name, address, pin string
	// Job details
	workAddress, company, position string
	salary                         int
}

// PersonBuilder struct
type PersonBuilder struct {
	person *Person
}

func (p *PersonBuilder) Build() *Person {
	return p.person
}

func NewPersonBuilder(name string) *PersonBuilder {
	return &PersonBuilder{person: &Person{name: name}}
}

// PersonAddressBuilder facet of PersonBuilder
type PersonAddressBuilder struct {
	PersonBuilder
}

func (p *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*p}
}

func (a *PersonAddressBuilder) At(location string) *PersonAddressBuilder {
	a.person.address = location
	return a
}

func (a *PersonAddressBuilder) WithPostalCode(code string) *PersonAddressBuilder {
	a.person.pin = code
	return a
}

// PersonJobBuilder facet of PersonBuilder
type PersonJobBuilder struct {
	PersonBuilder
}

func (p *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*p}
}

func (j *PersonJobBuilder) As(position string) *PersonJobBuilder {
	j.person.position = position
	return j
}

func (j *PersonJobBuilder) For(company string) *PersonJobBuilder {
	j.person.company = company
	return j
}

func (j *PersonJobBuilder) In(companyAddress string) *PersonJobBuilder {
	j.person.workAddress = companyAddress
	return j
}

func (j *PersonJobBuilder) WithSalary(salary int) *PersonJobBuilder {
	j.person.salary = salary
	return j
}

func RunBuilderFacet() {
	pb := NewPersonBuilder("Alex")
	pb.Lives().
		At("Bangalore").
		WithPostalCode("560102").
		Works().
		As("Software Engineer").
		For("IBM").
		In("Bangalore").
		WithSalary(150000)

	person := pb.Build()

	fmt.Println(*person)
}

func main() {
	RunBuilderFacet()
}
