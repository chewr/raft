package raft

//
// support for Raft and kvraft to save persistent
// Raft state (log &c) and k/v server snapshots.
//
type Persister interface {
	Copy() (Persister, error)
	SaveRaftState(data []byte) error
	ReadRaftState() ([]byte, error)
	RaftStateSize() int
	SaveSnapshot(snapshot []byte) error
	ReadSnapshot() ([]byte, error)
}
