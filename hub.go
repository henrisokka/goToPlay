package main

import "fmt"

type Hub struct {
	clients []int
}

var activeHub Hub = Hub{
	make([]int, 10),
}

func (h Hub) registerClient(id int) {
	h.clients = append(h.clients, id)
	fmt.Println(h.clients)
}

func newHub() {

	fmt.Println("Hub registering?")
}

func newClient(id int) {
	activeHub.registerClient(id)
}
