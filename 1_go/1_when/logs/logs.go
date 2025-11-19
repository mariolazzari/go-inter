package logs

import "time"

type Level byte

const (
	InfoLevel = iota + 1
	WarningLevel
	ErrorLevel
)

type Log struct {
	Time    time.Time
	Level   Level
	Message string
}

type DB interface {
	Query(start, end time.Time) ([]Log, error)
}

type Logs struct {
	db DB
}

func (ls *Logs) Query(start, end time.Time, level Level) ([]Log, error) {
	byDate, err := ls.db.Query(start, end)
	if err != nil {
		return nil, err
	}

	var logs []Log
	for _, l := range byDate {
		if l.Level >= level {
			logs = append(logs, l)
		}
	}

	return logs, nil
}
