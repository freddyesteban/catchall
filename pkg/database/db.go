package database

type ApiDatabase interface {
	Query() error
}