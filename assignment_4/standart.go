package assignment_4

import "fmt"

type StandartDelivery struct {
	transport Transport
}

func (s *StandartDelivery) Deliver() {
	fmt.Println("Standart delivery request")
	s.transport.DeliverItem()
}

func (s *StandartDelivery) ChooseTransport(transport Transport) {
	s.transport = transport
}
