package assignment_7

import (
	"software_design_patterns/assignment_7/observer"
	"software_design_patterns/assignment_7/subject"
)

func Print() {
	shirtItem := subject.NewItem("Nike Shirt")

	capItem := subject.NewItem("cap")

	customerFirst := &observer.Customer{Email: "abc@gmail.com"}
	customerSecond := &observer.Customer{Email: "xyz@gmail.com"}

	shirtItem.Register(customerFirst)
	shirtItem.Register(customerSecond)
	capItem.Register(customerFirst)

	shirtItem.UpdateAvailability()
	capItem.UpdateAvailability()

	capItem.Deregister(customerFirst)
	capItem.Register(customerSecond)
	capItem.UpdateAvailability()

}
