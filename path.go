package accessor

import (
	"bytes"
	"strings"
)

type Path interface {
	Key() string
	SubPath() (Path, bool)
	PushHead(key string) Path
}

type basicPath struct {
	key  string
	tail Path
}

func (p *basicPath) PushHead(key string) Path {
	return &basicPath{
		key,
		p,
	}
}

func (p *basicPath) Key() string {
	return p.key
}

func (p *basicPath) SubPath() (Path, bool) {
	if p.tail == nil {
		return nil, false
	}
	return p.tail, true
}

func (p *basicPath) String() string {
	buf := bytes.NewBufferString(p.key)
	tail, ok := p.SubPath()
	for ok {
		buf.WriteRune('/')
		buf.WriteString(tail.Key())
		tail, ok = tail.SubPath()
	}
	return buf.String()
}

type rootPath struct{}

func (p rootPath) PushHead(key string) Path {
	return &basicPath{
		key,
		nil,
	}
}

func (p rootPath) Key() string {
	return ""
}

func (p rootPath) SubPath() (Path, bool) {
	return nil, false
}

func (p rootPath) String() string {
	return "/"
}

func ParsePath(path string) (Path, error) {
	paths := strings.Split(strings.Trim(path, "/ "), "/")

	if len(paths) == 0 {
		return nil, NewInvalidPathError(path)
	}

	last := len(paths) - 1
	var p Path = rootPath{}
	for i := last; i >= 0; i-- {
		if paths[i] == "" {
			return nil, NewInvalidPathError(path)
		}
		p = p.PushHead(paths[i])
	}
	return p, nil
}
