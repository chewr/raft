package raft

type Endpoint interface {
	Call(request interface{}) (interface{}, error)
}

type Client interface {
	Call(serviceMethod string, args interface{}, reply interface{}) error
}
