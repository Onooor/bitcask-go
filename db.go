package bitcask_go

import (
	"bitcask-go/data"
	"bitcask-go/index"
	"sync"
)

type DB struct {
	options      Options
	mu           *sync.RWMutex
	activateFile *data.DataFile
	olderFiles   map[uint32]*data.DataFile
	index        index.Indexer
}

// put 写入 key/value 数据 key不能为空
func (db *DB) Put(key []byte, value []byte) error {
	if len(key) == 0 {
		return ErrKeyIsEmpty
	}
	logRecord := &data.LogRecord{
		Key:   key,
		Value: value,
		Type:  data.LogRecordNormal,
	}

	pos, err := db.appendLogRecord(logRecord)
	if err != nil {
		return err
	}

	if ok := db.index.Put(key, pos); !ok {
		return ErrIndexUpdateFailed
	}

	return nil

}

func (db *DB) Get(key []byte) ([]byte, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	if len(key) == 0 {
		return nil, ErrKeyIsEmpty
	}
	//从内存数据结构中取出key对应的索引信息

	logRecordPos := db.index.Get(key)

	if logRecordPos == nil {
		return nil, ErrKeyNotFound
	}
	var dataFile *data.DataFile
	if db.activateFile.FileId == logRecordPos.Fid {
		dataFile = db.activateFile
	} else {
		dataFile = db.olderFiles[logRecordPos.Fid]
	}
	if dataFile == nil {
		return nil, ErrDataFileNotFound
	}

	//根据偏移读取对应的数据
	logRecord, err := dataFile.ReadLogRecord(logRecordPos.Offset)
	if err != nil {
		return nil, err
	}

	if logRecord.Type != data.LogRecordNormal {
		return nil, ErrKeyNotFound
	}

	return logRecord.Value, nil
}

// 追加写入数据到活跃文件中
func (db *DB) appendLogRecord(logRecord *data.LogRecord) (*data.LogRecord, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	if db.activateFile == nil {
		if err := db.setActivateDataFile(); err != nil {
			return nil, err
		}
	}

	encRecord, size := data.EncodeLogRecord(logRecord)

	if db.activateFile.WriteOff+size > db.options.DataFileSize {
		if err := db.activateFile.Sync(); err != nil {
			return nil, err
		}

		db.olderFiles[db.activateFile.FileId] = db.activateFile

		if err := db.setActivateDataFile(); err != nil {
			return nil, err
		}

	}

	writeOff := db.activateFile.WriteOff
	if err := db.activateFile.Write(encRecord); err != nil {
		return nil, err
	}

	if db.options.SyncWrites {
		if err := db.activateFile.Sync(); err != nil {
			return nil, err
		}
	}

	pos := &data.LogRecordPos{Fid: db.activateFile.FileId, Offset: writeOff}
	return pos, nil

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
