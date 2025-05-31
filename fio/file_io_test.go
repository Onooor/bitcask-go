package fio

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestNewFileIOManager(t *testing.T) {
	fio, err := NewFileIOManager(filepath.Join("/Users/onee/Documents/kv-project/bitcask-go/tmp", "a.data"))

	assert.Nil(t, err)
	assert.NotNil(t, fio)

}

func TestFileIO_Write(t *testing.T) {
	fio, err := NewFileIOManager(filepath.Join("/Users/onee/Documents/kv-project/bitcask-go/tmp", "a.data"))

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
	fio, err := NewFileIOManager(filepath.Join("tmp", "a.data"))

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	//_, err := fio.Write([]byte("key-a\n"))
	//assert.Nil(t, err)

}
