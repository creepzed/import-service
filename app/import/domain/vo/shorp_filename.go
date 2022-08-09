package vo

import (
	"errors"
	"fmt"
)

type ShopFilename struct {
	value string
}

var (
	ErrShopFilenameInvalid = errors.New("ShopFilename is invalid")
)

func NewShopFilename(value string) (shopFilename ShopFilename, err error) {
	shopFilename = ShopFilename{value: value}
	if err = shopFilename.hasError(); err != nil {
		shopFilename = ShopFilename{}
	}
	return
}

func (s ShopFilename) hasError() error {
	if s.value == "" {
		return fmt.Errorf("%w: %s", ErrShopFilenameInvalid, s.value)
	}
	return nil
}

func (s ShopFilename) Value() string {
	return s.value
}
