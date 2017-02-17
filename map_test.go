package accessor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapAccessor_Get(t *testing.T) {
	type Input struct {
		Accessor Accessor
		Path     string
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
		Test{
			Title: "success",
			Input: Input{
				Accessor: MapAccessor(map[string]Accessor{
					"a": DummyAccessor{1},
				}),
				Path: "a",
			},
			Expect: Expect{
				Accessor: DummyAccessor{1},
				Err:      nil,
			},
		},
		Test{
			Title: "success nested",
			Input: Input{
				Accessor: MapAccessor(map[string]Accessor{
					"a": MapAccessor(map[string]Accessor{
						"b": MapAccessor(map[string]Accessor{
							"c": DummyAccessor{1},
						}),
					}),
				}),
				Path: "a/b/c",
			},
			Expect: Expect{
				Accessor: DummyAccessor{1},
				Err:      nil,
			},
		},
		Test{
			Title: "path error",
			Input: Input{
				Accessor: MapAccessor(map[string]Accessor{
					"a": DummyAccessor{1},
				}),
				Path: "x",
			},
			Expect: Expect{
				Accessor: nil,
				Err: &NoSuchPathError{
					Message: "no such key",
					Key:     "x",
					Stack:   nil,
				},
			},
		},
		Test{
			Title: "path error nested",
			Input: Input{
				Accessor: MapAccessor(map[string]Accessor{
					"a": MapAccessor(map[string]Accessor{
						"b": MapAccessor(map[string]Accessor{
							"c": DummyAccessor{1},
						}),
					}),
				}),
				Path: "a/b/x",
			},
			Expect: Expect{
				Accessor: nil,
				Err: &NoSuchPathError{
					Message: "no such key",
					Key:     "x",
					Stack:   []string{"b", "a"},
				},
			},
		},
	}

	for _, testCase := range table {
		t.Run(testCase.Title, func(t *testing.T) {
			assert := assert.New(t)

			path, err := ParsePath(testCase.Input.Path)
			assert.Nil(err)
			acc, err := testCase.Input.Accessor.Get(path)

			assert.Equal(testCase.Expect.Accessor, acc)
			assert.Equal(testCase.Expect.Err, err)
		})
	}
}

func TestMapAccessor_Set(t *testing.T) {
	type Input struct {
		Accessor Accessor
		Path     string
		BeSet    Accessor
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
		Test{
			Title: "success",
			Input: Input{
				Accessor: MapAccessor(map[string]Accessor{
					"a": DummyAccessor{1},
				}),
				Path:  "a",
				BeSet: DummyAccessor{2},
			},
			Expect: Expect{
				Accessor: MapAccessor(map[string]Accessor{
					"a": DummyAccessor{2},
				}),
				Err: nil,
			},
		},
		Test{
			Title: "success nested",
			Input: Input{
				Accessor: MapAccessor(map[string]Accessor{
					"a": MapAccessor(map[string]Accessor{
						"b": MapAccessor(map[string]Accessor{
							"c": DummyAccessor{1},
						}),
					}),
				}),
				Path:  "a/b",
				BeSet: DummyAccessor{2},
			},
			Expect: Expect{
				Accessor: MapAccessor(map[string]Accessor{
					"a": MapAccessor(map[string]Accessor{
						"b": DummyAccessor{2},
					}),
				}),
				Err: nil,
			},
		},
		Test{
			Title: "path error",
			Input: Input{
				Accessor: MapAccessor(map[string]Accessor{
					"a": DummyAccessor{1},
				}),
				Path:  "x",
				BeSet: DummyAccessor{2},
			},
			Expect: Expect{
				Accessor: MapAccessor(map[string]Accessor{
					"a": DummyAccessor{1},
				}),
				Err: &NoSuchPathError{
					Message: "no such key",
					Key:     "x",
					Stack:   nil,
				},
			},
		},
		Test{
			Title: "path error nested",
			Input: Input{
				Accessor: MapAccessor(map[string]Accessor{
					"a": MapAccessor(map[string]Accessor{
						"b": MapAccessor(map[string]Accessor{
							"c": DummyAccessor{1},
						}),
					}),
				}),
				Path:  "a/b/x",
				BeSet: DummyAccessor{2},
			},
			Expect: Expect{
				Accessor: MapAccessor(map[string]Accessor{
					"a": MapAccessor(map[string]Accessor{
						"b": MapAccessor(map[string]Accessor{
							"c": DummyAccessor{1},
						}),
					}),
				}),
				Err: &NoSuchPathError{
					Message: "no such key",
					Key:     "x",
					Stack:   []string{"b", "a"},
				},
			},
		},
	}

	for _, testCase := range table {
		t.Run(testCase.Title, func(t *testing.T) {
			assert := assert.New(t)

			path, err := ParsePath(testCase.Input.Path)
			assert.Nil(err)
			acc := testCase.Input.Accessor
			err = acc.Set(path, testCase.Input.BeSet)

			assert.Equal(testCase.Expect.Accessor, acc)
			assert.Equal(testCase.Expect.Err, err)
		})
	}
}

func TestMapAccessor_Unwrap(t *testing.T) {
	type Input struct {
		Accessor Accessor
	}
	type Expect struct {
		Accessor interface{}
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
				Accessor: MapAccessor(map[string]Accessor{
					"a": DummyAccessor{1},
				}),
			},
			Expect: Expect{
				Accessor: map[string]interface{}{
					"a": 1,
				},
			},
		},
		Test{
			Title: "success nested",
			Input: Input{
				Accessor: MapAccessor(map[string]Accessor{
					"a": MapAccessor(map[string]Accessor{
						"b": MapAccessor(map[string]Accessor{
							"c": DummyAccessor{1},
						}),
					}),
				}),
			},
			Expect: Expect{
				Accessor: map[string]interface{}{
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

			assert.Equal(testCase.Expect.Accessor, testCase.Input.Accessor.Unwrap())
		})
	}
}
