package mvcc

type Timestamp uint64

type Transaction struct {
	ID      Timestamp
	ReadTS  Timestamp
	WriteTS Timestamp
}

type Manager interface {
	Begin() (*Transaction, error)
	Commit(tx *Transaction) error
	Abort(tx *Transaction) error
}
