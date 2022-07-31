package vo

import (
	"errors"
	"fmt"
)

type ImportId struct {
	value string
}

var (
	ErrInvalidImportId = errors.New("importId is invalid")
)

func NewImportId(value string) (importId ImportId, err error) {
	importId = ImportId{value: value}
	if err = importId.hasError(); err != nil {
		importId = ImportId{}
	}
	return
}

func (i *ImportId) hasError() error {
	if i.value == "" {
		return fmt.Errorf("%w: %s", ErrInvalidImportId, i.value)
	}
	return nil
}

func (i ImportId) Value() string {
	return i.value
}
