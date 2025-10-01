package assignment_4

import "fmt"

func Print() {
	truck := &Truck{}
	drone := &Drone{}

	expressDelivery := &ExpressDelivery{}

	expressDelivery.ChooseTransport(truck)
	expressDelivery.Deliver()
	fmt.Println()

	expressDelivery.ChooseTransport(drone)
	expressDelivery.Deliver()
	fmt.Println()

	standartDelivery := &StandartDelivery{}

	standartDelivery.ChooseTransport(truck)
	standartDelivery.Deliver()
	fmt.Println()

	standartDelivery.ChooseTransport(drone)
	standartDelivery.Deliver()
	fmt.Println()
}
