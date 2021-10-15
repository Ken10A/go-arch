package main

import "fmt"

type Originator struct {
	state string
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) SetState(state string) {
	fmt.Println("Setting state to " + state)
	o.state = state
}

func (o *Originator) GetMemento() Memento {
	return Memento{o.state}
}

func (o *Originator) Restore(memento Memento) {
	o.state = memento.GetState()
}

type Memento struct {
	serializedState string
}

func (m *Memento) GetState() string {
	return m.serializedState
}

func Caretaker() {
	theOriginator := Originator{"A"}
	theOriginator.SetState("A")
	fmt.Println("theOriginator state = ", theOriginator.GetState())

	theMemento := theOriginator.GetMemento()

	theOriginator.SetState("unclean")
	fmt.Println("theOriginator state = ", theOriginator.GetState())

	theOriginator.Restore(theMemento)
	fmt.Println("RESTORED : theOriginator state = ", theOriginator.GetState())
}

func main() {
	Caretaker()
}
