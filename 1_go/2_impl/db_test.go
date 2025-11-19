package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOpen(t *testing.T) {
	_, err := Open("localhost")
	require.NoError(t, err)
}
