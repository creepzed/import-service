package vo

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

var (
	ErrInvalidID = errors.New("invalid ID")
)

type Id struct {
	value string
}

func NewId(value string) (id Id, err error) {
	id = Id{value: value}
	if err = id.hasError(); err != nil {
		return Id{}, err
	}
	return
}

func GenerateId() Id {
	id, _ := NewId(uuid.NewString())
	return id
}

func (i Id) hasError() error {
	_, err := uuid.Parse(i.value)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidID, i.value)
	}
	return nil
}

func (i Id) Value() string {
	return i.value
}
