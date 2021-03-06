package main

import (
	"errors"
	"fmt"
)

type Report interface {
	Execute()
}

type ConcreteReportA struct {
	receiver *Receiver
}

func (c *ConcreteReportA) Execute() {
	c.receiver.Action("ReportA")
}

type ConcreteReportB struct {
	receiver *Receiver
}

func (c *ConcreteReportB) Execute() {
	c.receiver.Action("ReportB")
}

type Receiver struct{}

func (r *Receiver) Action(msg string) {
	fmt.Println(msg)
}

type Invoker struct {
	repository []Report
}

func (i *Invoker) Schedule(cmd Report) {
	i.repository = append(i.repository, cmd)
}

func (i *Invoker) Run() {
	for _, cmd := range i.repository {
		cmd.Execute()
	}
}

type ChainedReceiver struct {
	canHandle string
	next      *ChainedReceiver
}

func (r *ChainedReceiver) SetNext(next *ChainedReceiver) {
	r.next = next
}

func (r *ChainedReceiver) Finish() error {
	fmt.Println(r.canHandle, " Receiver Finishing")
	return nil
}

func (r *ChainedReceiver) Handle(what string) error {
	if what == r.canHandle {
		return r.Finish()
	} else if r.next != nil {
		return r.next.Handle(what)
	} else {
		fmt.Println("No receiver could handle the request!")
		return errors.New("No receiver to handle")
	}
}

func main() {
	receiver := new(Receiver)
	reportA := &ConcreteReportA{receiver}
	reportB := &ConcreteReportB{receiver}
	invoker := new(Invoker)
	invoker.Schedule(reportA)
	invoker.Run()
	invoker.Schedule(reportB)
	invoker.Run()
}
