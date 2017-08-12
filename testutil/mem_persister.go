package testutil

import "sync"
import "github.com/chewr/raft/persistance"

//
// implementation of a simple in-memory persister
// suitable for testing only
type simplePersisterImpl struct {
	mu        sync.Mutex
	raftstate []byte
	snapshot  []byte
}

func NewSimplePersister() *simplePersisterImpl {
	return &simplePersisterImpl{}
}

func (ps *simplePersisterImpl) Copy() (persistance.Persister, error) {
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
