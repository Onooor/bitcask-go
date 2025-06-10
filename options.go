package bitcask_go

type Options struct {
	DirPath string

	DataFileSize int64

	SyncWrites bool

	IndexType IndexerType
}

type IndexerType = int8

const (
	BTreeIndex IndexerType = iota
	ART
)
