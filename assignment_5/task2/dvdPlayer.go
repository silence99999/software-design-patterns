package task2

import "fmt"

type DVDPlayer struct {
}

func (d *DVDPlayer) turnOnDVD() {
	fmt.Println("DVDPlayer turn on")
}

func (d *DVDPlayer) insertTheDisk(movie string) {
	fmt.Println("Inserting the Disk with ", movie)
}

func (d *DVDPlayer) PlayingTheMovie(movie string) {
	fmt.Println("Playing the movie: ", movie)
}

func (d *DVDPlayer) turnOffDVD() {
	fmt.Println("DVDPlayer turn off")
}
