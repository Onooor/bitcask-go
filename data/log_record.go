package data

type LogRecordType = byte

const (
	LogRecordNormal LogRecordType = iota
	LogRecordDeleted
)

type LogRecord struct {
	Key   []byte
	Value []byte
	Type  LogRecordType
}

type LogRecordPos struct {
	Fid    uint32 //文件id
	Offset int64
}

func EncodeLogRecord(record *LogRecord) ([]byte, int64) {
	return nil, 0
}
