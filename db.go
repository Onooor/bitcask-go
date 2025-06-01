package bitcask_go

import (
	"bitcask-go/data"
	"sync"
)

type DB struct {
	options      Options
	mu           *sync.RWMutex
	activateFile *data.DataFile
	olderFiles   map[uint32]*data.DataFile
}

func (db *DB) Put(key []byte, value []byte) error {
	if len(key) == 0 {
		return ErrKeyIsEmpty
	}
	log_record := data.LogRecord{
		Key:   key,
		Value: value,
		Type:  data.LogRecordNormal,
	}
}

func (db *DB) appendLogRecord(logRecord *data.LogRecord) (*data.LogRecord, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	if db.activateFile == nil {
		if err := db.setActivateDataFile(); err != nil {
			return nil, err
		}
	}

	encRecord, size := data.EncodeLogRecord(logRecord)

}

func (db *DB) setActivateDataFile() error {
	var initialFileId uint32 = 0
	if db.activateFile != nil {
		initialFileId = db.activateFile.FileId + 1
	}
	dataFile, err := data.OpenDataFile(db.options.DirPath, initialFileId)
	if err != nil {
		return err
	}
	db.activateFile = dataFile

}
