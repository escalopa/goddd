package entity

import "github.com/google/uuid"

type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}

func (i *Item) SetID(id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	i.ID = uuid
	return nil
}

func (i *Item) SetName(name string) error {
	i.Name = name
	return nil
}

func (i *Item) SetDescription(description string) error {
	i.Description = description
	return nil

}
