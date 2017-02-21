package accessor

import (
	"bytes"
	"fmt"
	"strings"
)

// Path is a sequence of object keys.
type Path interface {
	fmt.Stringer

	// Key returns a head key of the sequence.
	Key() string

	// SubPath returns a path excluding the beginning.
	SubPath() (Path, bool)

	// PushKey add a key to the head of the sequence.
	PushKey(key string) Path
}

type basicPath struct {
	key  string
	tail Path
}

func (p *basicPath) PushKey(key string) Path {
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

// PhantomPath is the Path without a key.
var PhantomPath Path = phantomPath{}

type phantomPath struct{}

func (p phantomPath) PushKey(key string) Path {
	return &basicPath{
		key,
		nil,
	}
}

func (p phantomPath) Key() string {
	return "???"
}

func (p phantomPath) SubPath() (Path, bool) {
	return nil, false
}

func (p phantomPath) String() string {
	return "???"
}

// ParsePath creates a Path from a slash(/)-separeted-keys.
func ParsePath(path string) (Path, error) {
	keys := strings.Split(strings.Trim(path, "/ "), "/")

	return NewPath(keys)
}

// NewPath creates a Path from keys.
func NewPath(keys []string) (Path, error) {
	if len(keys) == 0 {
		return nil, NewInvalidPathError("path is empty")
	}

	last := len(keys) - 1
	p := PhantomPath
	for i := last; i >= 0; i-- {
		if keys[i] == "" {
			return nil, NewInvalidPathError("empty key found")
		}
		p = p.PushKey(keys[i])
	}
	return p, nil
}

func newPath(keys ...string) Path {
	p, err := NewPath(keys)
	if err != nil {
		panic(err)
	}
	return p
}
