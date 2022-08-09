package domain

import (
	"bitbucket.org/ripleyx/import-service/app/import/domain/vo"
	"bitbucket.org/ripleyx/import-service/app/shared/domain/event"
	sharedVO "bitbucket.org/ripleyx/import-service/app/shared/domain/vo"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrRecordsAreInvalid = errors.New("the record set is invalid. One header record and at least one body record is required")
	ErrFileIsInvalid     = errors.New("the file is not valid, it needs to have more than two rows")
)

type Import struct {
	id               sharedVO.Id
	shopId           vo.ShopId
	externalImportId vo.ExternalImportId
	shopFilename     vo.ShopFilename
	headers          []vo.Head
	rows             [][]vo.Cell
	events           []event.Event
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

	anImportId, err := vo.NewExternalImportId(importIdRune)
	if err != nil {
		return Import{}, err
	}

	anShopFilename, err := vo.NewShopFilename(importFilename)
	if err != nil {
		return Import{}, err
	}

	anImport := Import{
		id:               sharedVO.GenerateId(),
		shopId:           aShopId,
		externalImportId: anImportId,
		shopFilename:     anShopFilename,
		events:           make([]event.Event, 0),
	}

	//TODO: create event

	return anImport, nil
}

func (i Import) Id() sharedVO.Id {
	return i.id
}

func (i Import) ShopId() vo.ShopId {
	return i.shopId
}

func (i Import) ExternalImportId() vo.ExternalImportId {
	return i.externalImportId
}

func (i Import) ShopFilename() vo.ShopFilename {
	return i.shopFilename
}

func (i Import) Headers() []vo.Head {
	return i.headers
}

func (i Import) Rows() [][]vo.Cell {
	return i.rows
}

func (i Import) AddRecordSet(records [][]string) error {
	if len(records) <= 1 {
		return fmt.Errorf("%w: contains %d row(s)", ErrRecordsAreInvalid, len(records))
	}
	for j, v := range records {
		if j == 0 {
			if err := i.AddHeader(v); err != nil {
				return err
			}
		} else {
			i.AddRow(v)
		}
	}
	return nil
}

func (i Import) AddHeader(record []string) error {
	headers := make([]vo.Head, 0)
	for _, col := range record {
		head, err := vo.NewHead(col)
		if err != nil {
			return err
		}
		headers = append(headers, head)
	}
	i.headers = headers
	return nil
}

func (i Import) AddRow(record []string) error {
	row := make([]vo.Cell, 0)
	for _, col := range record {
		head, err := vo.NewCell(col)
		if err != nil {
			return err
		}
		row = append(row, head)
	}
	i.rows = append(i.rows, row)
	return nil
}

func (i Import) Products() []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	for _, row := range i.rows {
		e := make(map[string]interface{})
		for j, col := range i.headers {
			e[col.Value()] = row[j].Value()
		}
		result = append(result, e)
	}
	return result
}
