package repository

import "context"

type GetObject interface {
	GetObject(ctx context.Context, objectName string) ([]byte, error)
}
