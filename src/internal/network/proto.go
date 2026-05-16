package network

type OpCode byte

const (
	OpGet    OpCode = 1
	OpSet    OpCode = 2
	OpDelete OpCode = 3
	OpQuery  OpCode = 4
)

type Header struct {
	Magic   [4]byte
	Version byte
	Op      OpCode
	Flags   byte
	Length  uint32
}
