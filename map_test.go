package accessor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapAccessor_Get(t *testing.T) {
	type Input struct {
		Accessor Accessor
		Path     string
		Paths    []string
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
				Paths: nil,
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
				Path:  "a",
				Paths: []string{"b", "c"},
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
				Path:  "x",
				Paths: nil,
			},
			Expect: Expect{
				Accessor: nil,
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
				Accessor: MapAccessor(map[string]Accessor{
					"a": MapAccessor(map[string]Accessor{
						"b": MapAccessor(map[string]Accessor{
							"c": DummyAccessor{1},
						}),
					}),
				}),
				Path:  "a",
				Paths: []string{"b", "x"},
			},
			Expect: Expect{
				Accessor: nil,
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

			acc, err := testCase.Input.Accessor.Get(testCase.Input.Path, testCase.Input.Paths...)

			assert.Equal(testCase.Expect.Accessor, acc)
			assert.Equal(testCase.Expect.Err, err)
		})
	}
}

func TestMapAccessor_Set(t *testing.T) {
	type Input struct {
		Accessor Accessor
		Path     string
		Paths    []string
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
				Paths: nil,
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
				Path:  "a",
				Paths: []string{"b"},
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
				Paths: nil,
				BeSet: DummyAccessor{2},
			},
			Expect: Expect{
				Accessor: MapAccessor(map[string]Accessor{
					"a": DummyAccessor{1},
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
				Accessor: MapAccessor(map[string]Accessor{
					"a": MapAccessor(map[string]Accessor{
						"b": MapAccessor(map[string]Accessor{
							"c": DummyAccessor{1},
						}),
					}),
				}),
				Path:  "a",
				Paths: []string{"b", "x"},
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
					Path:    "x",
					Stack:   []string{"b", "a"},
				},
			},
		},
	}

	for _, testCase := range table {
		t.Run(testCase.Title, func(t *testing.T) {
			assert := assert.New(t)

			acc := testCase.Input.Accessor
			err := acc.Set(testCase.Input.BeSet, testCase.Input.Path, testCase.Input.Paths...)

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
