package accessor

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValueAccessor_Get(t *testing.T) {
	type Input struct {
		Value interface{}
	}
	type Expect struct {
		ErrorMessage string
	}
	type Test struct {
		Title  string
		Input  Input
		Expect Expect
	}

	table := []Test{
		Test{
			Title:  "int",
			Input:  Input{1},
			Expect: Expect{"int(1) has no key"},
		},
		Test{
			Title:  "float",
			Input:  Input{1.2},
			Expect: Expect{"float64(1.2) has no key"},
		},
		Test{
			Title:  "string",
			Input:  Input{"hello"},
			Expect: Expect{"string(hello) has no key"},
		},
		Test{
			Title:  "bool",
			Input:  Input{true},
			Expect: Expect{"bool(true) has no key"},
		},
		Test{
			Title:  "time.Time",
			Input:  Input{time.Date(1992, 6, 18, 12, 34, 56, 00, time.UTC)},
			Expect: Expect{"time.Time(1992-06-18 12:34:56 +0000 UTC) has no key"},
		},
		Test{
			Title:  "nil",
			Input:  Input{nil},
			Expect: Expect{"<nil>(<nil>) has no key"},
		},
	}

	for _, testCase := range table {
		t.Run(testCase.Title, func(t *testing.T) {
			assert := assert.New(t)

			path, err := ParsePath("a")
			assert.Nil(err)
			bt := ValueAccessor{testCase.Input.Value}
			a, err := bt.Get(path)

			assert.Nil(a)
			assert.Equal(NewNoSuchPathError(testCase.Expect.ErrorMessage, "a"), err)
		})
	}
}

func TestValueAccessor_Set(t *testing.T) {
	type Input struct {
		Value interface{}
	}
	type Expect struct {
		ErrorMessage string
	}
	type Test struct {
		Title  string
		Input  Input
		Expect Expect
	}

	table := []Test{
		Test{
			Title:  "int",
			Input:  Input{1},
			Expect: Expect{"int(1) has no key"},
		},
		Test{
			Title:  "float",
			Input:  Input{1.2},
			Expect: Expect{"float64(1.2) has no key"},
		},
		Test{
			Title:  "string",
			Input:  Input{"hello"},
			Expect: Expect{"string(hello) has no key"},
		},
		Test{
			Title:  "bool",
			Input:  Input{true},
			Expect: Expect{"bool(true) has no key"},
		},
		Test{
			Title:  "time.Time",
			Input:  Input{time.Date(1992, 6, 18, 12, 34, 56, 00, time.UTC)},
			Expect: Expect{"time.Time(1992-06-18 12:34:56 +0000 UTC) has no key"},
		},
		Test{
			Title:  "nil",
			Input:  Input{nil},
			Expect: Expect{"<nil>(<nil>) has no key"},
		},
	}

	for _, testCase := range table {
		t.Run(testCase.Title, func(t *testing.T) {
			assert := assert.New(t)

			path, err := ParsePath("a")
			assert.Nil(err)
			bt := ValueAccessor{testCase.Input.Value}
			err = bt.Set(path, DummyAccessor{1})

			assert.Equal(bt.Value, testCase.Input.Value)
			assert.Equal(NewNoSuchPathError(testCase.Expect.ErrorMessage, "a"), err)
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
		Test{
			Title:  "int",
			Input:  Input{1},
			Expect: Expect{1},
		},
		Test{
			Title:  "float",
			Input:  Input{1.2},
			Expect: Expect{1.2},
		},
		Test{
			Title:  "string",
			Input:  Input{"hello"},
			Expect: Expect{"hello"},
		},
		Test{
			Title:  "bool",
			Input:  Input{true},
			Expect: Expect{true},
		},
		Test{
			Title:  "time.Time",
			Input:  Input{time.Date(1992, 6, 18, 12, 34, 56, 00, time.UTC)},
			Expect: Expect{time.Date(1992, 6, 18, 12, 34, 56, 00, time.UTC)},
		},
		Test{
			Title:  "nil",
			Input:  Input{nil},
			Expect: Expect{nil},
		},
	}

	for _, testCase := range table {
		t.Run(testCase.Title, func(t *testing.T) {
			assert := assert.New(t)

			assert.Equal(testCase.Expect.Value, ValueAccessor{testCase.Input.Value}.Unwrap())
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
		Test{
			Title: "success",
			Input: Input{
				Accessor:     ValueAccessor{1},
				ReturnsError: false,
			},
			Expect: Expect{
				Paths:        []string{"?"},
				ReturnsError: false,
			},
		},
		Test{
			Title: "error",
			Input: Input{
				Accessor:     ValueAccessor{1},
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
