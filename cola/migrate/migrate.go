package migrate

type migrater interface {
	migrate() error
	create() error
}
