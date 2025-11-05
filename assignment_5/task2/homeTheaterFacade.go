package task2

import "fmt"

type HomeTheaterFacade struct {
	dvd       *DVDPlayer
	projector *Projector
	sound     *SoundSystem
}

func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{
		dvd:       &DVDPlayer{},
		projector: &Projector{},
		sound:     &SoundSystem{},
	}
}

func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("Get ready to watch a movie")
	h.projector.On()
	h.sound.turnOnSoundSystem()
	h.sound.SetVolume(10)
	h.dvd.turnOnDVD()
	h.dvd.PlayingTheMovie(movie)
}

func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("Shutting movie theater down")
	h.dvd.turnOffDVD()
	h.sound.turnOffSoundSystem()
	h.projector.Off()
}
