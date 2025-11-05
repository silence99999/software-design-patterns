package assignment_6

type CourierDelivery struct {
}

func (c CourierDelivery) CalculateCost(distance int) int {
	costPerKM := 10
	courierDeliveryCost := distance * costPerKM
	return courierDeliveryCost
}
