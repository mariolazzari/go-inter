package rotate

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotator(t *testing.T) {
	rootPath := t.TempDir()
	out, err := New(rootPath, 100)
	require.NoError(t, err, "New")

	logger := log.New(out, "[test] ", log.LstdFlags)

	for i := 0; i < 5; i++ {
		logger.Printf("info: Go Rocks!")
	}

	out.Close()

	matches, err := filepath.Glob(filepath.Join(rootPath, "log*.txt"))
	require.NoError(t, err, "glob")
	require.Equal(t, len(matches), 2)
}
