package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type State struct {
	sounds []SoundInfo
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

func (s *State) handleAction(a Action) {
	fmt.Println("Before append:")
	fmt.Println(s.sounds)
	s.sounds = append(s.sounds, a.Sound)
}
