package assignment_4

import "fmt"

type ExpressDelivery struct {
	transport Transport
}

func (e *ExpressDelivery) Deliver() {
	fmt.Println("Express delivery request")
	e.transport.DeliverItem()
}

func (e *ExpressDelivery) ChooseTransport(transport Transport) {
	e.transport = transport
}
