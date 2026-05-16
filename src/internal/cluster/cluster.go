package cluster

type NodeID uint64

type Config struct {
	NodeID     NodeID
	Peers      []NodeID
	ListenAddr string
	DataDir    string
}

type Node interface {
	Start() error
	Stop() error
	IsLeader() bool
	Leader() NodeID
}
