package vo

import (
	"errors"
	"fmt"
)

type ProductsFile struct {
	value string
}

var (
	ErrProductsFileInvalid = errors.New("ProductsFile is invalid")
)

func NewProductsFile(value string) (productsFile ProductsFile, err error) {
	productsFile = ProductsFile{value: value}
	if productsFile.hasError(); err != nil {
		productsFile = ProductsFile{}
	}
	return
}

func (f ProductsFile) hasError() error {
	if f.value == "" {
		return fmt.Errorf("%w: %s", ErrProductsFileInvalid, f.value)
	}
	return nil
}
