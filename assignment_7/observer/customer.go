package observer

import "fmt"

type Customer struct {
	Email string
}

func (c *Customer) Update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.Email, itemName)
}

func (c *Customer) GetEmail() string {
	return c.Email
}
