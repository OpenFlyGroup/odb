package object

type TypeID uint32

type Object struct {
	Type    TypeID
	Version uint32
	Flags   uint32
	Data    []byte
}

type Codec interface {
	Encode(obj *Object) ([]byte, error)
	Decode(data []byte) (*Object, error)
}
