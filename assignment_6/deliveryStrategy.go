package assignment_6

type DeliveryStrategy interface {
	CalculateCost(distance int) int
}
