package entity

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID
	Name       string
	Age        uint8
	Email      string
	IsVerified bool
}

func (c User) ValidateData() error {
	if c.Age == 0 {
		return errors.New("Age Required")
	}
	if c.Email == "" {
		return errors.New("Email Required")
	}
	if c.Name == "" {
		return errors.New("Name Required")
	}
	return nil
}

func NewUser(dto User) *User {
	return &User{
		ID:    uuid.New(),
		Name:  dto.Name,
		Age:   dto.Age,
		Email: dto.Email,
	}
}
