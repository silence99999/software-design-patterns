package assignment_6

import "fmt"

func Print() {
	order := Order{
		Distance: 20,
	}

	order.Strategy = CourierDelivery{}
	fmt.Println("Courier delivery cost: ", order.GetDeliveryCost())

	order.Strategy = PickupDelivery{}
	fmt.Println("Pickup delivery cost: ", order.GetDeliveryCost())

	order.Strategy = PostDelivery{}
	fmt.Println("Post delivery cost: ", order.GetDeliveryCost())
}
