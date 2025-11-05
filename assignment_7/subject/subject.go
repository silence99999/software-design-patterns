package subject

import "software_design_patterns/assignment_7/observer"

type Subject interface {
	register(observer observer.IObserver)
	deregister(observer observer.IObserver)
	notifyAll()
}
