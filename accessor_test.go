package accessor

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewAccessor(t *testing.T) {
	type Input struct {
		Value interface{}
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
			Title: "map[string]interface{}",
			Input: Input{
				Value: map[string]interface{}{
					"int": 1,
				},
			},
			Expect: Expect{
				Accessor: MapAccessor(map[string]Accessor{
					"int": ValueAccessor{1},
				}),
				Err: nil,
			},
		},
		{
			Title: "map[interface{}]interface{}",
			Input: Input{
				Value: map[interface{}]interface{}{
					"int": 1,
				},
			},
			Expect: Expect{
				Accessor: MapAccessor(map[string]Accessor{
					"int": ValueAccessor{1},
				}),
				Err: nil,
			},
		},
		{
			Title: "[]interface{}",
			Input: Input{
				Value: []interface{}{
					1,
					"string",
					true,
				},
			},
			Expect: Expect{
				Accessor: SliceAccessor([]Accessor{
					ValueAccessor{1},
					ValueAccessor{"string"},
					ValueAccessor{true},
				}),
				Err: nil,
			},
		},
		{
			Title: "basic type",
			Input: Input{
				Value: time.Date(1992, 6, 18, 12, 34, 56, 78, time.UTC),
			},
			Expect: Expect{
				Accessor: ValueAccessor{time.Date(1992, 6, 18, 12, 34, 56, 78, time.UTC)},
				Err:      nil,
			},
		},
		{
			Title: "complex",
			Input: Input{
				Value: map[string]interface{}{
					"name": "me",
					"age":  18,
					"friends": []interface{}{
						map[string]interface{}{
							"name": "hello",
						},
						map[string]interface{}{
							"name": "world",
						},
					},
					"birthday": time.Date(1992, 6, 18, 12, 34, 56, 78, time.UTC),
					"nickname": nil,
				},
			},
			Expect: Expect{
				Accessor: MapAccessor(map[string]Accessor{
					"name": ValueAccessor{"me"},
					"age":  ValueAccessor{18},
					"friends": SliceAccessor([]Accessor{
						MapAccessor(map[string]Accessor{
							"name": ValueAccessor{"hello"},
						}),
						MapAccessor(map[string]Accessor{
							"name": ValueAccessor{"world"},
						}),
					}),
					"birthday": ValueAccessor{time.Date(1992, 6, 18, 12, 34, 56, 78, time.UTC)},
					"nickname": ValueAccessor{nil},
				}),
				Err: nil,
			},
		},
	}

	for _, testCase := range table {
		t.Run(testCase.Title, func(t *testing.T) {
			assert := assert.New(t)

			acc, err := NewAccessor(testCase.Input.Value)

			assert.Equal(testCase.Expect.Accessor, acc)
			assert.Equal(testCase.Expect.Err, err)
		})
	}
}
