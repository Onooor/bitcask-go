package data

import (
	"bitcask-go/fio"
)

type DataFile struct {
	FileId    uint32
	WriteOff  int64
	IoManager fio.IOManager
}

func OpenDataFile(dirPath string, filedId uint32) (*DataFile, error) {
	return nil, nil
}
