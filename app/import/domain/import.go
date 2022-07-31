package domain

import (
	"bitbucket.org/ripleyx/import-service/app/import/domain/vo"
	"bitbucket.org/ripleyx/import-service/app/shared/application/event"
	"strings"
)

type Import struct {
	shopId   vo.ShopId
	importId vo.ImportId
	shopFile vo.ShopFilename

	events []event.Event
}

const (
	patter = "-_."
)

func NewImport(importFilename string) (Import, error) {

	splitFunc := func(r rune) bool {
		return strings.ContainsRune(patter, r)
	}
	runes := strings.FieldsFunc(importFilename, splitFunc)

	shopIdRune := runes[0]
	importIdRune := runes[1]

	aShopId, err := vo.NewShopId(shopIdRune)
	if err != nil {
		return Import{}, err
	}

	anImportId, err := vo.NewImportId(importIdRune)
	if err != nil {
		return Import{}, err
	}

	anShopFilename, err := vo.NewShopFilename(importFilename)
	if err != nil {
		return Import{}, err
	}

	anImport := Import{
		shopId:   aShopId,
		importId: anImportId,
		shopFile: anShopFilename,
		events:   make([]event.Event, 0),
	}

	return anImport, nil
}
