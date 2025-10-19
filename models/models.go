package models

type Auth struct {
	username string
	password string
}

type Connection struct {
	service        string
	lifetime       int
	source_address int
}

type Network struct {
	bandwith    int
	subnet      int
	connections []Connection
}

type Service struct {
	name string
	port int
}

// Attacks

// Defense
type Firewall struct {
	level int
}

// Progression
type Progress struct {
	power     int
	resources int
	money     int
}

// Main factory
type Factory struct {
	network  Network
	firewall Firewall
	services []Service
	progress Progress
}
