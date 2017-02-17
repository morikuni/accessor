package accessor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceAccessor_Get(t *testing.T) {
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
				Accessor: SliceAccessor([]Accessor{
					DummyAccessor{1},
				}),
				Path:  "0",
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
				Accessor: SliceAccessor([]Accessor{
					SliceAccessor([]Accessor{
						SliceAccessor([]Accessor{
							DummyAccessor{1},
						}),
					}),
				}),
				Path:  "0",
				Paths: []string{"0", "0"},
			},
			Expect: Expect{
				Accessor: DummyAccessor{1},
				Err:      nil,
			},
		},
		Test{
			Title: "not a number",
			Input: Input{
				Accessor: SliceAccessor([]Accessor{
					DummyAccessor{1},
				}),
				Path:  "x",
				Paths: nil,
			},
			Expect: Expect{
				Accessor: nil,
				Err: &NoSuchPathError{
					Message: "not a number",
					Path:    "x",
					Stack:   nil,
				},
			},
		},
		Test{
			Title: "index out of range",
			Input: Input{
				Accessor: SliceAccessor([]Accessor{
					DummyAccessor{1},
				}),
				Path:  "1",
				Paths: nil,
			},
			Expect: Expect{
				Accessor: nil,
				Err: &NoSuchPathError{
					Message: "index out of range",
					Path:    "1",
					Stack:   nil,
				},
			},
		},
		Test{
			Title: "path error nested",
			Input: Input{
				Accessor: SliceAccessor([]Accessor{
					SliceAccessor([]Accessor{
						SliceAccessor([]Accessor{
							DummyAccessor{1},
						}),
					}),
				}),
				Path:  "0",
				Paths: []string{"0", "1"},
			},
			Expect: Expect{
				Accessor: nil,
				Err: &NoSuchPathError{
					Message: "index out of range",
					Path:    "1",
					Stack:   []string{"0", "0"},
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

func TestSliceAccessor_Set(t *testing.T) {
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
				Accessor: SliceAccessor([]Accessor{
					DummyAccessor{1},
				}),
				Path:  "0",
				Paths: nil,
				BeSet: DummyAccessor{2},
			},
			Expect: Expect{
				Accessor: SliceAccessor([]Accessor{
					DummyAccessor{2},
				}),
				Err: nil,
			},
		},
		Test{
			Title: "success nested",
			Input: Input{
				Accessor: SliceAccessor([]Accessor{
					SliceAccessor([]Accessor{
						SliceAccessor([]Accessor{
							DummyAccessor{1},
						}),
					}),
				}),
				Path:  "0",
				Paths: []string{"0"},
				BeSet: DummyAccessor{2},
			},
			Expect: Expect{
				Accessor: SliceAccessor([]Accessor{
					SliceAccessor([]Accessor{
						DummyAccessor{2},
					}),
				}),
				Err: nil,
			},
		},
		Test{
			Title: "not a number",
			Input: Input{
				Accessor: SliceAccessor([]Accessor{
					DummyAccessor{1},
				}),
				Path:  "x",
				Paths: nil,
				BeSet: DummyAccessor{2},
			},
			Expect: Expect{
				Accessor: SliceAccessor([]Accessor{
					DummyAccessor{1},
				}),
				Err: &NoSuchPathError{
					Message: "not a number",
					Path:    "x",
					Stack:   nil,
				},
			},
		},
		Test{
			Title: "index out of range",
			Input: Input{
				Accessor: SliceAccessor([]Accessor{
					DummyAccessor{1},
				}),
				Path:  "1",
				Paths: nil,
				BeSet: DummyAccessor{2},
			},
			Expect: Expect{
				Accessor: SliceAccessor([]Accessor{
					DummyAccessor{1},
				}),
				Err: &NoSuchPathError{
					Message: "index out of range",
					Path:    "1",
					Stack:   nil,
				},
			},
		},
		Test{
			Title: "path error nested",
			Input: Input{
				Accessor: SliceAccessor([]Accessor{
					SliceAccessor([]Accessor{
						SliceAccessor([]Accessor{
							DummyAccessor{1},
						}),
					}),
				}),
				Path:  "0",
				Paths: []string{"0", "1"},
				BeSet: DummyAccessor{2},
			},
			Expect: Expect{
				Accessor: SliceAccessor([]Accessor{
					SliceAccessor([]Accessor{
						SliceAccessor([]Accessor{
							DummyAccessor{1},
						}),
					}),
				}),
				Err: &NoSuchPathError{
					Message: "index out of range",
					Path:    "1",
					Stack:   []string{"0", "0"},
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

func TestSliceAccessor_Unwrap(t *testing.T) {
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
				Accessor: SliceAccessor([]Accessor{
					DummyAccessor{1},
				}),
			},
			Expect: Expect{
				Accessor: []interface{}{1},
			},
		},
		Test{
			Title: "success nested",
			Input: Input{
				Accessor: SliceAccessor([]Accessor{
					SliceAccessor([]Accessor{
						SliceAccessor([]Accessor{
							DummyAccessor{1},
						}),
					}),
				}),
			},
			Expect: Expect{
				Accessor: []interface{}{
					[]interface{}{
						[]interface{}{
							1,
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
