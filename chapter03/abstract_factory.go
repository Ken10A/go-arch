package main

import "fmt"

type Reservation interface{}
type Invoice interface{}

type AbstractFactory interface {
	CreateReservation() Reservation
	CreateInvoidce() Invoice
}

type HotelFactory struct{}

func (f HotelFactory) CreateReservation() Reservation {
	return new(HotelReservation)
}

func (f HotelFactory) CreateInvoidce() Invoice {
	return new(HotelInvoice)
}

type FlightFactory struct{}

func (f FlightFactory) CreateReservation() Reservation {
	return new(FlightReservation)
}

func (f FlightFactory) CreateInvoidce() Invoice {
	return new(FlightInvoice)
}

func GetFactory(vertical string) AbstractFactory {
	var factory AbstractFactory
	switch vertical {
	case "flight":
		factory = FlightFactory{}
	case "hotel":
		factory = HotelFactory{}
	}
	return factory
}

type HotelReservation struct{}
type HotelInvoice struct{}
type FlightReservation struct{}
type FlightInvoice struct{}

func main() {
	hotelFactory := GetFactory("hotel")
	reservation := hotelFactory.CreateReservation()
	invoice := hotelFactory.CreateInvoidce()

	fmt.Printf("%T\n", reservation)
	fmt.Printf("%T\n", invoice)
}
