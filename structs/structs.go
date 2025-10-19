package models

type Connection struct {
	port     int
	service  string
	lifetime int
}

type Network struct {
	bandwith    int
	subnet      int
	connections []Connection
}

type Factory struct {
	network Network
}




