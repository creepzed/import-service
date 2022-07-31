package split

import (
	"bitbucket.org/ripleyx/import-service/app/import/domain/repository"
	"bitbucket.org/ripleyx/import-service/app/shared/infrastructure/log"
	"bytes"
	"context"
	"encoding/csv"
)

type SplitService interface {
	Do(ctx context.Context, command ImportSplitCommand) error
}

type splitService struct {
	repository repository.GetObject
}

func NewSplitService(repositoryGetObject repository.GetObject) *splitService {
	return &splitService{
		repository: repositoryGetObject,
	}
}

func (g splitService) Do(ctx context.Context, command ImportSplitCommand) error {
	file, err := g.repository.GetObject(ctx, command.Filename())
	if err != nil {
		return err
	}

	csvReader := csv.NewReader(bytes.NewBuffer(file))

	data, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	log.Debug("+v", data)
	return nil
}
