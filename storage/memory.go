package storage

import (
	"fmt"

	"github.com/IanRivas/go-crud/model"
)

// Memory en lugar de usar una base de datos por ahora
type Memory struct {
	currentID int
	Persons   map[int]model.Person
}

// NewMemory constructor
func NewMemory() Memory {
	persons := make(map[int]model.Person)

	return Memory{
		currentID: 0,
		Persons:   persons,
	}
}

// Create para agregar el objeto memory , la nueva persona
func (m *Memory) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	m.currentID++
	m.Persons[m.currentID] = *person

	return nil
}

// Update actualiza una persona en el slice de memory
func (m *Memory) Update(ID int, person *model.Person) error {
	if person != nil {
		return model.ErrPersonCanNotBeNil
	}

	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	m.Persons[ID] = *person

	return nil
}

// Delete borra de la memoria la persona
func (m *Memory) Delete(ID int) error {
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}
	delete(m.Persons, ID)

	return nil
}
