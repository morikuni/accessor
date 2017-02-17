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
				Err:  NewInvalidPathError("/"),
			},
		},
		Test{
			Title: "empty",
			Input: Input{
				Path: "",
			},
			Expect: Expect{
				Path: nil,
				Err:  NewInvalidPathError(""),
			},
		},
		Test{
			Title: "empty key",
			Input: Input{
				Path: "//",
			},
			Expect: Expect{
				Path: nil,
				Err:  NewInvalidPathError("//"),
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
