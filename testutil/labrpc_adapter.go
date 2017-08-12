package testutil

import "fmt"
import "github.com/chewr/6.824-2016/labrpc"
import "github.com/chewr/raft"

type labrpcAdapter struct {
	e *labrpc.ClientEnd
}

func NewLabRpcAdapter(e *labrpc.ClientEnd) raft.Client {
	return &labrpcAdapter{
		e: e,
	}
}

func (c *labrpcAdapter) Call(serviceMethod string, args interface{}, reply interface{}) error {
	var err error
	if c.e.Call(serviceMethod, args, reply) {
		err = fmt.Errorf("Call %s failed", serviceMethod)
	}
	return err
}
