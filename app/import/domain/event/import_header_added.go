package event

import (
	sharedEvent "bitbucket.org/ripleyx/import-service/app/shared/domain/event"
)

const ImportHeaderAddedType sharedEvent.Type = "event.import.header.added"

type ImportHeaderAddedEvent struct {
	ShopId   string   `json:"shop_id"`
	ImportId string   `json:"import_id"`
	ShopFile string   `json:"shop_file"`
	Headers  []string `json:"headers"`
	sharedEvent.BaseEvent
}

func NewImportHeaderAddedEvent(shopId, importId, shopFile string, headers []string) ImportHeaderAddedEvent {
	return ImportHeaderAddedEvent{
		ShopId:   shopId,
		ImportId: importId,
		ShopFile: shopFile,
		Headers:  headers,
	}
}
