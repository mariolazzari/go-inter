package logs

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type mockDB struct{}

func (m mockDB) Query(start, end time.Time) ([]Log, error) {
	return nil, fmt.Errorf("no connection")
}

func TestLogs_QueryError(t *testing.T) {
	ls := Logs{db: mockDB{}}
	end := time.Now()
	start := end.Add(-24 * time.Hour)
	_, err := ls.Query(start, end, InfoLevel)
	require.Error(t, err)
}
