package accessor

import (
	"bytes"
)

type Path interface {
	Key() string
	SubPath() (Path, bool)
	PushHead(key string) Path
}

type path struct {
	key  string
	tail *path
}

func (p *path) PushHead(key string) Path {
	return &path{
		key,
		p,
	}
}

func (p *path) Key() string {
	return p.key
}

func (p *path) SubPath() (Path, bool) {
	if p.tail == nil {
		return nil, false
	}
	return p.tail, true
}

func (p *path) String() string {
	buf := bytes.NewBufferString(p.key)
	tail := p.tail
	for tail != nil {
		buf.WriteRune('/')
		buf.WriteString(tail.key)
		tail = tail.tail
	}
	return buf.String()
}

type emptyPath struct{}

func (p emptyPath) PushHead(key string) Path {
	return &path{
		key,
		nil,
	}
}

func (p emptyPath) Key() string {
	return ""
}

func (p emptyPath) SubPath() (Path, bool) {
	return nil, false
}

func (p emptyPath) String() string {
	return "?"
}
