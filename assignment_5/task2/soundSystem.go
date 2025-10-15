package task2

import "fmt"

type SoundSystem struct{}

func (s *SoundSystem) turnOnSoundSystem() {
	fmt.Println("Sound System is ON")
}

func (s *SoundSystem) SetVolume(level int) {
	fmt.Println("Volume set to", level)
}

func (s *SoundSystem) turnOffSoundSystem() {
	fmt.Println("Sound System is OFF")
}
