package accessor

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewObject(t *testing.T) {
	type Input struct {
		Value interface{}
	}
	type Expect struct {
		Object Object
		Err    error
	}
	type Test struct {
		Title  string
		Input  Input
		Expect Expect
	}

	table := []Test{
		Test{
			Title: "map[string]interface{}",
			Input: Input{
				Value: map[string]interface{}{
					"int": 1,
				},
			},
			Expect: Expect{
				Object: MapObject(map[string]Object{
					"int": ValueObject{1},
				}),
				Err: nil,
			},
		},
		Test{
			Title: "map[interface{}]interface{}",
			Input: Input{
				Value: map[interface{}]interface{}{
					"int": 1,
				},
			},
			Expect: Expect{
				Object: MapObject(map[string]Object{
					"int": ValueObject{1},
				}),
				Err: nil,
			},
		},
		Test{
			Title: "[]interface{}",
			Input: Input{
				Value: []interface{}{
					1,
					"string",
					true,
				},
			},
			Expect: Expect{
				Object: SliceObject([]Object{
					ValueObject{1},
					ValueObject{"string"},
					ValueObject{true},
				}),
				Err: nil,
			},
		},
		Test{
			Title: "basic type",
			Input: Input{
				Value: time.Date(1992, 6, 18, 12, 34, 56, 78, time.UTC),
			},
			Expect: Expect{
				Object: ValueObject{time.Date(1992, 6, 18, 12, 34, 56, 78, time.UTC)},
				Err:    nil,
			},
		},
		Test{
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
				Object: MapObject(map[string]Object{
					"name": ValueObject{"me"},
					"age":  ValueObject{18},
					"friends": SliceObject([]Object{
						MapObject(map[string]Object{
							"name": ValueObject{"hello"},
						}),
						MapObject(map[string]Object{
							"name": ValueObject{"world"},
						}),
					}),
					"birthday": ValueObject{time.Date(1992, 6, 18, 12, 34, 56, 78, time.UTC)},
					"nickname": ValueObject{nil},
				}),
				Err: nil,
			},
		},
	}

	for _, testCase := range table {
		t.Run(testCase.Title, func(t *testing.T) {
			assert := assert.New(t)

			obj, err := NewObject(testCase.Input.Value)

			assert.Equal(testCase.Expect.Object, obj)
			assert.Equal(testCase.Expect.Err, err)
		})
	}
}
