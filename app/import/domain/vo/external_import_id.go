package vo

import (
	"errors"
	"fmt"
)

type ExternalImportId struct {
	value string
}

var (
	ErrInvalidExternalImportId = errors.New("external importId is invalid")
)

func NewExternalImportId(value string) (externalImportId ExternalImportId, err error) {
	externalImportId = ExternalImportId{value: value}
	if err = externalImportId.hasError(); err != nil {
		externalImportId = ExternalImportId{}
	}
	return
}

func (i ExternalImportId) hasError() error {
	if i.value == "" {
		return fmt.Errorf("%w: %s", ErrInvalidExternalImportId, i.value)
	}
	return nil
}

func (i ExternalImportId) Value() string {
	return i.value
}
