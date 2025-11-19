package main

type Record struct {
	Key  uint
	Data []byte
}

// Puller pulls a single record from the database.
type Puller interface {
	Pull(r *Record) error
}
