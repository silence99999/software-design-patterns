package subject

import (
	"fmt"
	"software_design_patterns/assignment_7/observer"
)

type Item struct {
	observerList []observer.IObserver
	name         string
	inStock      bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}
func (i *Item) UpdateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *Item) Register(o observer.IObserver) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) Deregister(o observer.IObserver) {
	i.observerList = removeFromslice(i.observerList, o)
}

func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.Update(i.name)
	}
}

func removeFromslice(observerList []observer.IObserver, observerToRemove observer.IObserver) []observer.IObserver {
	for i, observer := range observerList {
		if observerToRemove.GetEmail() == observer.GetEmail() {
			observerList = append(observerList[:i], observerList[i+1:]...)
			return observerList
		}
	}
	return observerList
}
