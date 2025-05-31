package fio

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func destroyFile(name string) {
	if err := os.Remove(name); err != nil {
		panic(err)
	}
}

func TestNewFileIOManager(t *testing.T) {

	path := filepath.Join("../tmp", "001.data")
	fio, err := NewFileIOManager(path)
	defer destroyFile(path)

	assert.Nil(t, err)
	assert.NotNil(t, fio)

}

func TestFileIO_Write(t *testing.T) {
	path := filepath.Join("../tmp", "001.data")
	fio, err := NewFileIOManager(path)
	defer destroyFile(path)

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	n, err := fio.Write([]byte("hello world1\n"))
	assert.Equal(t, 13, n)
	assert.Nil(t, err)

	n, err = fio.Write([]byte("hello goland2\n"))
	assert.Equal(t, 14, n)
	t.Log(n, err)
	n, err = fio.Write([]byte("hello clion3\n"))
	assert.Equal(t, 13, n)
	t.Log(n, err)

}

func TestFileIO_Read(t *testing.T) {
	path := filepath.Join("../tmp", "001.data")
	fio, err := NewFileIOManager(path)
	defer destroyFile(path)

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	_, err = fio.Write([]byte("key-a\n"))
	assert.Nil(t, err)

	_, err = fio.Write([]byte("key-b\n"))
	assert.Nil(t, err)

	b := make([]byte, 6)
	n, err := fio.Read(b, 0)
	t.Log(b, n)
	assert.Equal(t, 6, n)
	assert.Equal(t, []byte("key-a\n"), b)

	b = make([]byte, 6)
	n, err = fio.Read(b, 6)
	t.Log(b, n)
	assert.Equal(t, 6, n)
	assert.Equal(t, []byte("key-b\n"), b)
	t.Log(string(b))

}

func TestFileIO_Sync(t *testing.T) {
	path := filepath.Join("../tmp", "001.data")
	fio, err := NewFileIOManager(path)
	defer destroyFile(path)

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	err = fio.Sync()
	assert.Nil(t, err)
}

func TestFileIO_Close(t *testing.T) {
	path := filepath.Join("../tmp", "001.data")
	fio, err := NewFileIOManager(path)
	defer destroyFile(path)

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	err = fio.Close()
	assert.Nil(t, err)
}
