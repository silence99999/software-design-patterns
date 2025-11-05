package assignment_6

type PostDelivery struct {
}

func (p PostDelivery) CalculateCost(distance int) int {
	costPerKM := 5
	postDeliveryCost := distance * costPerKM
	return postDeliveryCost
}
