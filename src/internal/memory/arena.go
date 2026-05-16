package memory

type Allocator interface {
	Alloc(size int) ([]byte, error)
	Free(ptr []byte) error
	Reset()
}
