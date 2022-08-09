package vo

import (
	"errors"
	"fmt"
)

type ShopId struct {
	value string
}

var (
	ErrShopIdInvalid = errors.New("ShopId is invalid")
)

func NewShopId(value string) (shopId ShopId, err error) {
	shopId = ShopId{value: value}
	if err = shopId.hasError(); err != nil {
		shopId = ShopId{}
	}
	return
}

func (s ShopId) hasError() error {
	if s.value == "" {
		return fmt.Errorf("%w: %s", ErrShopIdInvalid, s.value)
	}
	return nil
}

func (s ShopId) Value() string {
	return s.value
}
