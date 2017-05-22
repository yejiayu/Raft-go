package internal

import "net"

type state uint

const (
	leader state = iota
	follower
	candidate
)

// Raft raft.
type Raft struct {
	currentTerm uint
	state       state
	voteCount   uint
	commitIndex string
	conn        net.Conn
}
