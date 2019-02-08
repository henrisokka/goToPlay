package main

import "fmt"

type state struct {
	sounds []soundInfo
}

type synth interface {
}

type soundInfo struct {
	Vel    int
	Freq   int
	Length int
}

type action struct {
	Type  string
	Sound soundInfo
}

func actionHandler(ie incomingEvent) {
	switch ie.a.Type {
	case "START_SOUND":
		startSound(ie)
	case "STOP_SOUND":
		stopSound(ie)
	default:
		fmt.Println("Can't detect the type")
	}
}

func startSound(ie incomingEvent) {
	fmt.Println("Start")
	fmt.Println(ie.a.Sound)
}

func stopSound(ie incomingEvent) {
	fmt.Println("Stop")
}

func (s *state) newSound(si soundInfo) {
	s.sounds = append(s.sounds, si)
}
