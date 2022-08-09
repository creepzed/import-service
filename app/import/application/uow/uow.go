package uow

type UnitOfWork interface {
	CreateTransaction() error
	Commit() error
	Rollback() error
	Save() error
}
