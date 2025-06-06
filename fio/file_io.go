package fio

import "os"

type FileIO struct {
	fd *os.File
}

func NewFileIOManager(filename string) (*FileIO, error) {
	fd, err := os.OpenFile(
		filename,
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		DataFileRerm,
	)
	if err != nil {
		return nil, err
	}
	return &FileIO{fd: fd}, nil
}

func (fio *FileIO) Read(b []byte, offset int64) (int, error) {
	return fio.fd.ReadAt(b, offset)
}

func (fio *FileIO) Write(b []byte) (int, error) {
	return fio.fd.Write(b)
}

func (fio *FileIO) Sync() error {
	return fio.fd.Sync()
}

func (fio *FileIO) Close() error {
	return fio.fd.Close()
}
