package model

// Community estructura de una comunidad
type Community struct {
	Name string
}

// Communities slice de comunidades
type Communities []Community

// Person estructura de una persona
type Person struct {
	Name        string
	Age         uint8
	Communities Communities
}

// Persons slice de personas
type Persons []Person
