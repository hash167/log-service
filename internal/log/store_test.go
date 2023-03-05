package log

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	write = []byte("hello world")
	width = uint(len(write)) + lenWidth
)

func TestStoreAppendRead(t *testing.T) {
	f, err := os.CreateTemp("", "sample")
	require.NoError(t, err)
	defer os.Remove(f.Name())
	defer f.Close()
	s, err := newStore(f)
	require.NoError(t, err)
	testAppend(t, s)
}

func testAppend(t *testing.T, s *store) {
	t.Helper()
	for i := uint64(0); i < 4; i++ {
		n, pos, err := s.Append(write)
		require.NoError(t, err)
		require.Equal(t, width, n)
		require.Equal(t, i*uint64(width), pos)
	}
}
