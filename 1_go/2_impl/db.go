package db

import (
	"fmt"
)

type DBError struct {
	Address string
	Reason  string
}

func (e *DBError) Error() string {
	return fmt.Sprintf("%s (address=%q)", e.Reason, e.Address)
}

type DB struct{}

func Open(host string) (*DB, error) {
	var err *DBError

	// TODO: Connect
	return nil, err
}
