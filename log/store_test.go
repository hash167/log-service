package log

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	write = []byte("hello world")
	width = uint64(len(write)) + lenWidth
)

func TestStoreAppendRead(t *testing.T) {
	f, err := os.CreateTemp("", "sample")
	require.NoError(t, err)
	defer os.Remove(f.Name())
	defer f.Close()
	s, err := NewStore(f)
	require.NoError(t, err)
	testAppend(t, s)
	testRead(t, s)
	testReadAt(t, s)
}

func testAppend(t *testing.T, s *store) {
	t.Helper()
	for i := uint64(1); i < 4; i++ {
		n, pos, err := s.Append(write)
		require.NoError(t, err)
		require.Equal(t, pos+n, i*width)
	}
}

func testRead(t *testing.T, s *store) {
	t.Helper()
	var pos uint64
	for i := uint64(1); i < 4; i++ {
		data, err := s.Read(pos)
		require.NoError(t, err)
		require.Equal(t, write, data)
		pos += width
	}
}

func testReadAt(t *testing.T, s *store) {
	t.Helper()
	for i, off := uint64(1), int64(0); i < 4; i++ {
		data := make([]byte, lenWidth)
		n, err := s.ReadAt(data, off)
		require.NoError(t, err)
		require.Equal(t, lenWidth, n)
		off += int64(n)
		size := enc.Uint64(data)
		data = make([]byte, size)
		n, err = s.ReadAt(data, off)
		require.NoError(t, err)
		require.Equal(t, size, uint64(n))
		require.Equal(t, int(size), n)
		off += int64(n)
	}

}

func TestStoreClose(t *testing.T) {
	f, err := os.CreateTemp("", "sample2")
	require.NoError(t, err)
	defer os.Remove(f.Name())
	defer f.Close()
	s, err := NewStore(f)
	require.NoError(t, err)
	s.Append(write)
	_, beforeSize, err := openFile(f.Name())
	require.NoError(t, err)
	err = s.Close()
	require.NoError(t, err)
	_, afterSize, err := openFile(f.Name())
	require.NoError(t, err)
	require.True(t, beforeSize < afterSize)
}

func openFile(name string) (file *os.File, size int64, err error) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, 0, err
	}
	fi, err := f.Stat()
	if err != nil {
		return nil, 0, err
	}
	return f, fi.Size(), nil
}
