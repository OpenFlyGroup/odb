package storage

type WAL interface {
	Append(entry LogEntry) (int64, error)
	Read(from int64) ([]LogEntry, error)
	Truncate(to int64) error
	Close() error
}

type LogEntry struct {
	SeqNum   int64
	Op       OpType
	Key      []byte
	Value    []byte
	Checksum uint32
}

type OpType byte

const (
	OpSet    OpType = 1
	OpDelete OpType = 2
)
