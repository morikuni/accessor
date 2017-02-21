package accessor

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValueAccessor_Get(t *testing.T) {
	type Input struct {
		Accessor Accessor
		Path     Path
	}
	type Expect struct {
		Accessor Accessor
		Err      error
	}
	type Test struct {
		Title  string
		Input  Input
		Expect Expect
	}

	table := []Test{
		{
			Title: "error",
			Input: Input{
				Accessor: &ValueAccessor{1},
				Path:     newPath("a"),
			},
			Expect: Expect{
				Accessor: nil,
				Err:      NewNoSuchPathError("int(1) has no key", "a"),
			},
		},
		{
			Title: "phantom path",
			Input: Input{
				Accessor: &ValueAccessor{1},
				Path:     thePhantomPath,
			},
			Expect: Expect{
				Accessor: &ValueAccessor{1},
				Err:      nil,
			},
		},
	}

	for _, testCase := range table {
		t.Run(testCase.Title, func(t *testing.T) {
			assert := assert.New(t)

			acc := testCase.Input.Accessor
			value, err := acc.Get(testCase.Input.Path)

			assert.Equal(testCase.Expect.Accessor, value)
			assert.Equal(testCase.Expect.Err, err)
		})
	}
}

func TestValueAccessor_Set(t *testing.T) {
	type Input struct {
		Accessor Accessor
		Path     Path
		BeSet    interface{}
	}
	type Expect struct {
		Accessor Accessor
		Err      error
	}
	type Test struct {
		Title  string
		Input  Input
		Expect Expect
	}

	table := []Test{
		{
			Title: "error",
			Input: Input{
				Accessor: &ValueAccessor{1},
				Path:     newPath("a"),
				BeSet:    2,
			},
			Expect: Expect{
				Accessor: &ValueAccessor{1},
				Err:      NewNoSuchPathError("int(1) has no key", "a"),
			},
		},
		{
			Title: "phantom path",
			Input: Input{
				Accessor: &ValueAccessor{1},
				Path:     thePhantomPath,
				BeSet:    2,
			},
			Expect: Expect{
				Accessor: &ValueAccessor{2},
				Err:      nil,
			},
		},
	}

	for _, testCase := range table {
		t.Run(testCase.Title, func(t *testing.T) {
			assert := assert.New(t)

			acc := testCase.Input.Accessor
			err := acc.Set(testCase.Input.Path, testCase.Input.BeSet)

			assert.Equal(testCase.Expect.Accessor, acc)
			assert.Equal(testCase.Expect.Err, err)
		})
	}
}

func TestValueAccessor_Unwrap(t *testing.T) {
	type Input struct {
		Value interface{}
	}
	type Expect struct {
		Value interface{}
	}
	type Test struct {
		Title  string
		Input  Input
		Expect Expect
	}

	table := []Test{
		{
			Title:  "int",
			Input:  Input{1},
			Expect: Expect{1},
		},
		{
			Title:  "float",
			Input:  Input{1.2},
			Expect: Expect{1.2},
		},
		{
			Title:  "string",
			Input:  Input{"hello"},
			Expect: Expect{"hello"},
		},
		{
			Title:  "bool",
			Input:  Input{true},
			Expect: Expect{true},
		},
		{
			Title:  "time.Time",
			Input:  Input{time.Date(1992, 6, 18, 12, 34, 56, 00, time.UTC)},
			Expect: Expect{time.Date(1992, 6, 18, 12, 34, 56, 00, time.UTC)},
		},
		{
			Title:  "nil",
			Input:  Input{nil},
			Expect: Expect{nil},
		},
	}

	for _, testCase := range table {
		t.Run(testCase.Title, func(t *testing.T) {
			assert := assert.New(t)

			acc := &ValueAccessor{testCase.Input.Value}
			assert.Equal(testCase.Expect.Value, acc.Unwrap())
		})
	}
}

func TestValueAccessor_Foreach(t *testing.T) {
	type Input struct {
		Accessor     Accessor
		ReturnsError bool
	}
	type Expect struct {
		Paths        []string
		ReturnsError bool
	}
	type Test struct {
		Title  string
		Input  Input
		Expect Expect
	}

	table := []Test{
		{
			Title: "success",
			Input: Input{
				Accessor:     &ValueAccessor{1},
				ReturnsError: false,
			},
			Expect: Expect{
				Paths:        []string{"???"},
				ReturnsError: false,
			},
		},
		{
			Title: "error",
			Input: Input{
				Accessor:     &ValueAccessor{1},
				ReturnsError: true,
			},
			Expect: Expect{
				Paths:        []string{},
				ReturnsError: true,
			},
		},
	}

	testErr := errors.New("test error")
	for _, testCase := range table {
		t.Run(testCase.Title, func(t *testing.T) {
			assert := assert.New(t)

			expextedPaths := map[string]struct{}{}
			for _, p := range testCase.Expect.Paths {
				expextedPaths[p] = struct{}{}
			}
			var expextedErr error
			if testCase.Expect.ReturnsError {
				expextedErr = testErr
			}

			paths := map[string]struct{}{}
			err := testCase.Input.Accessor.Foreach(func(path Path, _ interface{}) error {
				if testCase.Input.ReturnsError {
					return testErr
				}
				paths[path.String()] = struct{}{}
				return nil
			})

			assert.Equal(expextedPaths, paths)
			assert.Equal(expextedErr, err)
		})
	}
}
