package vo

import (
	"errors"
	"fmt"
	"strings"
)

type Head struct {
	value string
}

var (
	ErrHeadIsInvalid = errors.New("head is invalid, cannot contain blank characters")
)

func NewHead(value string) (head Head, err error) {
	head = Head{value: value}
	if err = head.hasError(); err != nil {
		return Head{}, err
	}
	return head, nil
}

func (h Head) hasError() error {
	if strings.Contains(h.value, " ") {
		return fmt.Errorf("%w: %s", ErrHeadIsInvalid, h.value)
	}
	return nil
}

func (h Head) Value() string {
	return h.value
}
