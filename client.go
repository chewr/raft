package raft

import "fmt"
import "github.com/chewr/6.824-2016/labrpc"

type Endpoint interface {
	Call(request interface{}) (interface{}, error)
}

type Client interface {
	Call(serviceMethod string, args interface{}, reply interface{}) error
}

type labrpcAdapter struct {
	e *labrpc.ClientEnd
}

func NewLabRpcAdapter(e *labrpc.ClientEnd) Client {
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
