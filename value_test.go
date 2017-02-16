package undef

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestValueObject_Get(t *testing.T) {
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

			bt := ValueObject{testCase.Input.Value}
			o, err := bt.Get("a")

			assert.Nil(o)
			assert.Equal(&NoSuchPathError{
				Message: testCase.Expect.ErrorMessage,
				Path:    "a",
				Stack:   nil,
			}, err)
		})
	}
}

func TestValueObject_Set(t *testing.T) {
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

			bt := ValueObject{testCase.Input.Value}
			err := bt.Set(DummyObject{1}, "a")

			assert.Equal(bt.Value, testCase.Input.Value)
			assert.Equal(&NoSuchPathError{
				Message: testCase.Expect.ErrorMessage,
				Path:    "a",
				Stack:   nil,
			}, err)
		})
	}
}

func TestValueObject_Unwrap(t *testing.T) {
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

			assert.Equal(testCase.Expect.Value, ValueObject{testCase.Input.Value}.Unwrap())
		})
	}
}
