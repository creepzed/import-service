package postgres

import (
	"bitbucket.org/ripleyx/import-service/app/import/domain"
	"context"
	"database/sql"
	"github.com/google/uuid"
)

type importRepository struct {
	db *sql.DB
}

func NewImportRepository(db *sql.DB) *importRepository {
	return &importRepository{
		db: db,
	}
}

func (i *importRepository) Save(ctx context.Context, anImport domain.Import) error {

	err := i.insertImportHeader(ctx, anImport)
	if err != nil {
		return err
	}

	err = i.insertImportRows(ctx, anImport)
	if err != nil {
		return err
	}
	return nil
}

func (i importRepository) insertImportHeader(ctx context.Context, anImport domain.Import) error {
	stmt, err := i.db.Prepare("INSERT INTO import(id, version ,shop_id, external_import_id, headers) VALUES(?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.ExecContext(ctx, anImport.Id().Value(), 0, anImport.ShopId().Value(),
		anImport.ExternalImportId().Value(), anImport.Headers()); err != nil {
		return err
	}
	return nil
}

func (i importRepository) insertImportRows(ctx context.Context, anImport domain.Import) error {
	stmt, err := i.db.Prepare("INSERT INTO import(id, import_id, status, complete, json_data_import, json_data_error) VALUES(?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, v := range anImport.Rows() {
		if _, err := stmt.ExecContext(ctx, uuid.NewString(), anImport.Id().Value(), "", "", v, nil); err != nil {
			return err
		}
	}
	return nil
}
