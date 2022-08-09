package split

import (
	"bitbucket.org/ripleyx/import-service/app/import/domain"
	"bitbucket.org/ripleyx/import-service/app/import/domain/repository"
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
)

type ImportSplitService interface {
	Do(ctx context.Context, command ImportSplitCommand) error
}

type importSplitService struct {
	repositoryObject repository.GetObject
	repositoryImport repository.SaveImport
}

func NewImportSplitService(repositoryGetObject repository.GetObject, repositorySaveImport repository.SaveImport) *importSplitService {
	return &importSplitService{
		repositoryObject: repositoryGetObject,
		repositoryImport: repositorySaveImport,
	}
}

func (g importSplitService) Do(ctx context.Context, command ImportSplitCommand) error {
	anImport, err := domain.NewImport(command.filename)
	if err != nil {
		return err
	}

	//TODO: Validar que la importanción no exista
	bytesFile, err := g.repositoryObject.GetObject(ctx, command.Filename())
	if err != nil {
		return err
	}

	csvReader := csv.NewReader(bytes.NewReader(bytesFile))
	csvReader.LazyQuotes = true
	csvReader.Comma = ';'

	data, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	err = anImport.AddRecordSet(data)
	if err != nil {
		return err
	}
	/*
		err = g.repositoryImport.Save(ctx, anImport)
		if err != nil {
			return err
		}
	*/
	return fmt.Errorf("sin error")
}
