package assignment_6

type Order struct {
	Distance int
	Strategy DeliveryStrategy
}

func (o Order) GetDeliveryCost() int {
	return o.Strategy.CalculateCost(o.Distance)
}
