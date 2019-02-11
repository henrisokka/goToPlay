package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type State struct {
	Sounds []SoundWithPlayer
}

type Event struct {
	conn   *websocket.Conn
	action Action
}

type Action struct {
	Type  string
	Sound SoundInfo
}

type SoundInfo struct {
	Vel    int
	Freq   int
	Length int
}

type SoundWithPlayer struct {
	Sound  SoundInfo
	Player string
}

func (s *State) handleAction(a Action, clientID string) {
	switch a.Type {
	case "START_SOUND":
		s.startSound(a.Sound, clientID)
	case "STOP_SOUND":
		s.stopSound(a.Sound)
	default:
		fmt.Println("Can't detect action type")
	}
}

func (s *State) startSound(so SoundInfo, clientID string) {
	sound := SoundWithPlayer{so, clientID}
	s.Sounds = append(s.Sounds, sound)
}

func (s *State) stopSound(so SoundInfo) {
	fmt.Println("Stop sound")
}
