package task2

func Print() {
	homeTheater := NewHomeTheaterFacade()
	homeTheater.WatchMovie("1+1")
	homeTheater.EndMovie()
}
