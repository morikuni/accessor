package undef

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapObject_Get(t *testing.T) {
	type Input struct {
		Object Object
		Path   string
		Paths  []string
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
			Title: "success",
			Input: Input{
				Object: MapObject(map[string]Object{
					"a": DummyObject{1},
				}),
				Path:  "a",
				Paths: nil,
			},
			Expect: Expect{
				Object: DummyObject{1},
				Err:    nil,
			},
		},
		Test{
			Title: "success nested",
			Input: Input{
				Object: MapObject(map[string]Object{
					"a": MapObject(map[string]Object{
						"b": MapObject(map[string]Object{
							"c": DummyObject{1},
						}),
					}),
				}),
				Path:  "a",
				Paths: []string{"b", "c"},
			},
			Expect: Expect{
				Object: DummyObject{1},
				Err:    nil,
			},
		},
		Test{
			Title: "path error",
			Input: Input{
				Object: MapObject(map[string]Object{
					"a": DummyObject{1},
				}),
				Path:  "x",
				Paths: nil,
			},
			Expect: Expect{
				Object: nil,
				Err: &NoSuchPathError{
					Message: "no such key",
					Path:    "x",
					Stack:   nil,
				},
			},
		},
		Test{
			Title: "path error nested",
			Input: Input{
				Object: MapObject(map[string]Object{
					"a": MapObject(map[string]Object{
						"b": MapObject(map[string]Object{
							"c": DummyObject{1},
						}),
					}),
				}),
				Path:  "a",
				Paths: []string{"b", "x"},
			},
			Expect: Expect{
				Object: nil,
				Err: &NoSuchPathError{
					Message: "no such key",
					Path:    "x",
					Stack:   []string{"b", "a"},
				},
			},
		},
	}

	for _, testCase := range table {
		t.Run(testCase.Title, func(t *testing.T) {
			assert := assert.New(t)

			obj, err := testCase.Input.Object.Get(testCase.Input.Path, testCase.Input.Paths...)

			assert.Equal(testCase.Expect.Object, obj)
			assert.Equal(testCase.Expect.Err, err)
		})
	}
}

func TestMapObject_Set(t *testing.T) {
	type Input struct {
		Object Object
		Path   string
		Paths  []string
		BeSet  Object
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
			Title: "success",
			Input: Input{
				Object: MapObject(map[string]Object{
					"a": DummyObject{1},
				}),
				Path:  "a",
				Paths: nil,
				BeSet: DummyObject{2},
			},
			Expect: Expect{
				Object: MapObject(map[string]Object{
					"a": DummyObject{2},
				}),
				Err: nil,
			},
		},
		Test{
			Title: "success nested",
			Input: Input{
				Object: MapObject(map[string]Object{
					"a": MapObject(map[string]Object{
						"b": MapObject(map[string]Object{
							"c": DummyObject{1},
						}),
					}),
				}),
				Path:  "a",
				Paths: []string{"b"},
				BeSet: DummyObject{2},
			},
			Expect: Expect{
				Object: MapObject(map[string]Object{
					"a": MapObject(map[string]Object{
						"b": DummyObject{2},
					}),
				}),
				Err: nil,
			},
		},
		Test{
			Title: "path error",
			Input: Input{
				Object: MapObject(map[string]Object{
					"a": DummyObject{1},
				}),
				Path:  "x",
				Paths: nil,
				BeSet: DummyObject{2},
			},
			Expect: Expect{
				Object: MapObject(map[string]Object{
					"a": DummyObject{1},
				}),
				Err: &NoSuchPathError{
					Message: "no such key",
					Path:    "x",
					Stack:   nil,
				},
			},
		},
		Test{
			Title: "path error nested",
			Input: Input{
				Object: MapObject(map[string]Object{
					"a": MapObject(map[string]Object{
						"b": MapObject(map[string]Object{
							"c": DummyObject{1},
						}),
					}),
				}),
				Path:  "a",
				Paths: []string{"b", "x"},
				BeSet: DummyObject{2},
			},
			Expect: Expect{
				Object: MapObject(map[string]Object{
					"a": MapObject(map[string]Object{
						"b": MapObject(map[string]Object{
							"c": DummyObject{1},
						}),
					}),
				}),
				Err: &NoSuchPathError{
					Message: "no such key",
					Path:    "x",
					Stack:   []string{"b", "a"},
				},
			},
		},
	}

	for _, testCase := range table {
		t.Run(testCase.Title, func(t *testing.T) {
			assert := assert.New(t)

			obj := testCase.Input.Object
			err := obj.Set(testCase.Input.BeSet, testCase.Input.Path, testCase.Input.Paths...)

			assert.Equal(testCase.Expect.Object, obj)
			assert.Equal(testCase.Expect.Err, err)
		})
	}
}

func TestMapObject_Unwrap(t *testing.T) {
	type Input struct {
		Object Object
	}
	type Expect struct {
		Object interface{}
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
				Object: MapObject(map[string]Object{
					"a": DummyObject{1},
				}),
			},
			Expect: Expect{
				Object: map[string]interface{}{
					"a": 1,
				},
			},
		},
		Test{
			Title: "success nested",
			Input: Input{
				Object: MapObject(map[string]Object{
					"a": MapObject(map[string]Object{
						"b": MapObject(map[string]Object{
							"c": DummyObject{1},
						}),
					}),
				}),
			},
			Expect: Expect{
				Object: map[string]interface{}{
					"a": map[string]interface{}{
						"b": map[string]interface{}{
							"c": 1,
						},
					},
				},
			},
		},
	}

	for _, testCase := range table {
		t.Run(testCase.Title, func(t *testing.T) {
			assert := assert.New(t)

			assert.Equal(testCase.Expect.Object, testCase.Input.Object.Unwrap())
		})
	}
}
