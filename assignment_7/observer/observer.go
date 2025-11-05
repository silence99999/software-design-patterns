package observer

type IObserver interface {
	Update(string)
	GetEmail() string
}
