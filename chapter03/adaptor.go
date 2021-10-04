package main

import "fmt"

type Adaptee struct{}

func (a Adaptee) ExistingMethod() {
	fmt.Println("using existing method")
}

type Adaptor struct {
	adaptee *Adaptee
}

func NewAdaptor() *Adaptor {
	return &Adaptor{new(Adaptee)}
}

func (a *Adaptor) ExpectedMethod() {
	fmt.Println("doing some work")
	a.adaptee.ExistingMethod()
}

func main() {
	adaptor := NewAdaptor()
	adaptor.ExpectedMethod()
}
