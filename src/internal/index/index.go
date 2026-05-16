package index

type Index interface {
	Insert(key []byte, ptr uint64) error
	Delete(key []byte) error
	Lookup(key []byte) (uint64, error)
	Scan(start, end []byte) (Iterator, error)
	Close() error
}

type Iterator interface {
	Next() bool
	Key() []byte
	Value() uint64
	Close() error
}
