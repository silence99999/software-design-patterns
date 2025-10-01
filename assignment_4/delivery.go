package assignment_4

type DeliveryType interface {
	Deliver()
	ChooseTransport(Transport)
}
