package bitcask_go

import (
	"bitcask-go/data"
	"bitcask-go/index"
	"errors"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type DB struct {
	options      Options
	mu           *sync.RWMutex
	activateFile *data.DataFile
	olderFiles   map[uint32]*data.DataFile
	index        index.Indexer
}

func Open(options Options) (*DB, error) {
	if err := checkOptions(options); err != nil {
		return nil, err
	}

	if _, err := os.Stat(options.DirPath); os.IsNotExist(err) {
		if err := os.Mkdir(options.DirPath, os.ModePerm); err != nil {
			return nil, err
		}
	}

	//初始化 DB实例结构体
	db := &DB{
		options:    options,
		mu:         new(sync.RWMutex),
		olderFiles: make(map[uint32]*data.DataFile),
		index:      index.NewIndexer(options.IndexType),
	}

	if err := db.loadDataFiles(); err != nil {
		return nil, err
	}

	if err := db.loadIndexFromDataFiles(); err != nil {
		return nil, err
	}
}

// Put put 写入 key/value 数据 key不能为空
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
func (db *DB) appendLogRecord(logRecord *data.LogRecord) (*data.LogRecordPos, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	if db.activateFile == nil {
		if err := db.setActiveDataFile(); err != nil {
			return nil, err
		}
	}

	encRecord, size := data.EncodeLogRecord(logRecord)

	if db.activateFile.WriteOff+size > db.options.DataFileSize {
		if err := db.activateFile.Sync(); err != nil {
			return nil, err
		}

		db.olderFiles[db.activateFile.FileId] = db.activateFile

		if err := db.setActiveDataFile(); err != nil {
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

func (db *DB) setActiveDataFile() error {
	var initialFileId uint32 = 0
	if db.activateFile != nil {
		initialFileId = db.activateFile.FileId + 1
	}
	dataFile, err := data.OpenDataFile(db.options.DirPath, initialFileId)
	if err != nil {
		return err
	}
	db.activateFile = dataFile
	return nil
}

func (db *DB) loadDataFiles() error {
	dirEntries, err := os.ReadDir(db.options.DirPath)
	if err != nil {
		return err
	}
	var fileIds []int

	for _, entry := range dirEntries {
		if strings.HasSuffix(entry.Name(), data.DataFileNameSuffix) {
			splitNames := strings.Split(entry.Name(), ".")
			fileId, err := strconv.Atoi(splitNames[0])
			if err != nil {
				return ErrDataDirectoryCorrupted
			}
			fileIds = append(fileIds, fileId)
		}
	}
	sort.Ints(fileIds)

	for i, fid := range fileIds {
		dataFile, err := data.OpenDataFile(db.options.DirPath, uint32(fid))
		if err != nil {
			return err
		}
		if i == len(fileIds)-1 {
			db.activateFile = dataFile
		} else {
			db.olderFiles[uint32(fid)] = dataFile
		}
	}
	return nil
}

func (db *DB) loadIndexFromDataFiles() error {

}

func checkOptions(options Options) error {
	if options.DirPath == "" {
		return errors.New("database path is empty")
	}

	if options.DataFileSize <= 0 {
		return errors.New("database Data File Size must be grater than zero")
	}
	return nil
}
