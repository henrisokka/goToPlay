package main

import "fmt"

type state struct {
}

type synth interface {
}

type action struct {
	Type string
}

func actionListener(a action) {
	switch a.Type {
	case "START_SOUND":
		startSound(a)
	case "STOP_SOUND":
		stopSound(a)
	default:
		fmt.Println("Can't detect the type")
	}
}

func startSound(a action) {
	fmt.Println("Start")
}

func stopSound(a action) {
	fmt.Println("Stop")
}
