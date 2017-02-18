package accessor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsePath(t *testing.T) {
	type Input struct {
		Path string
	}
	type Expect struct {
		Path Path
		Err  error
	}
	type Test struct {
		Title  string
		Input  Input
		Expect Expect
	}

	table := []Test{
		Test{
			Title: "basic",
			Input: Input{
				Path: "a/b/0",
			},
			Expect: Expect{
				Path: &basicPath{"a", &basicPath{"b", &basicPath{"0", nil}}},
				Err:  nil,
			},
		},
		Test{
			Title: "many slash",
			Input: Input{
				Path: "/a/b/0/",
			},
			Expect: Expect{
				Path: &basicPath{"a", &basicPath{"b", &basicPath{"0", nil}}},
				Err:  nil,
			},
		},
		Test{
			Title: "only slash",
			Input: Input{
				Path: "/",
			},
			Expect: Expect{
				Path: nil,
				Err:  NewInvalidPathError("empty key found"),
			},
		},
		Test{
			Title: "empty",
			Input: Input{
				Path: "",
			},
			Expect: Expect{
				Path: nil,
				Err:  NewInvalidPathError("empty key found"),
			},
		},
		Test{
			Title: "empty key",
			Input: Input{
				Path: "//",
			},
			Expect: Expect{
				Path: nil,
				Err:  NewInvalidPathError("empty key found"),
			},
		},
	}

	for _, testCase := range table {
		t.Run(testCase.Title, func(t *testing.T) {
			assert := assert.New(t)

			path, err := ParsePath(testCase.Input.Path)

			assert.Equal(testCase.Expect.Path, path)
			assert.Equal(testCase.Expect.Err, err)
		})
	}
}

func TestNewPath(t *testing.T) {
	type Input struct {
		Keys []string
	}
	type Expect struct {
		Path Path
		Err  error
	}
	type Test struct {
		Title  string
		Input  Input
		Expect Expect
	}

	table := []Test{
		Test{
			Title: "basic",
			Input: Input{
				Keys: []string{"a", "b", "c"},
			},
			Expect: Expect{
				Path: &basicPath{"a", &basicPath{"b", &basicPath{"c", nil}}},
				Err:  nil,
			},
		},
		Test{
			Title: "include empty string",
			Input: Input{
				Keys: []string{"a", "", "c"},
			},
			Expect: Expect{
				Path: nil,
				Err:  NewInvalidPathError("empty key found"),
			},
		},
		Test{
			Title: "empty",
			Input: Input{
				Keys: nil,
			},
			Expect: Expect{
				Path: nil,
				Err:  NewInvalidPathError("path is empty"),
			},
		},
	}

	for _, testCase := range table {
		t.Run(testCase.Title, func(t *testing.T) {
			assert := assert.New(t)

			path, err := NewPath(testCase.Input.Keys)

			assert.Equal(testCase.Expect.Path, path)
			assert.Equal(testCase.Expect.Err, err)
		})
	}
}
