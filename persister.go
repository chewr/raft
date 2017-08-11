package raft

import "sync"

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

type simplePersisterImpl struct {
	mu        sync.Mutex
	raftstate []byte
	snapshot  []byte
}

func NewSimplePersister() Persister {
	return &simplePersisterImpl{}
}

func (ps *simplePersisterImpl) Copy() (Persister, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	np := &simplePersisterImpl{}
	np.raftstate = ps.raftstate
	np.snapshot = ps.snapshot
	return np, nil
}

func (ps *simplePersisterImpl) SaveRaftState(data []byte) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.raftstate = data
	return nil
}

func (ps *simplePersisterImpl) ReadRaftState() ([]byte, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	return ps.raftstate, nil
}

func (ps *simplePersisterImpl) RaftStateSize() int {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	return len(ps.raftstate)
}

func (ps *simplePersisterImpl) SaveSnapshot(snapshot []byte) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.snapshot = snapshot
	return nil
}

func (ps *simplePersisterImpl) ReadSnapshot() ([]byte, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	return ps.snapshot, nil
}
