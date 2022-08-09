package repository

import (
	"bitbucket.org/ripleyx/import-service/app/import/domain"
	"context"
)

type GetObject interface {
	GetObject(ctx context.Context, objectName string) ([]byte, error)
}

type SaveImport interface {
	Save(ctx context.Context, anImport domain.Import) error
}
