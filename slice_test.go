package undef

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceObject_Get(t *testing.T) {
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
				Object: SliceObject([]Object{
					DummyObject{1},
				}),
				Path:  "0",
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
				Object: SliceObject([]Object{
					SliceObject([]Object{
						SliceObject([]Object{
							DummyObject{1},
						}),
					}),
				}),
				Path:  "0",
				Paths: []string{"0", "0"},
			},
			Expect: Expect{
				Object: DummyObject{1},
				Err:    nil,
			},
		},
		Test{
			Title: "not a number",
			Input: Input{
				Object: SliceObject([]Object{
					DummyObject{1},
				}),
				Path:  "x",
				Paths: nil,
			},
			Expect: Expect{
				Object: nil,
				Err: &PathError{
					Message: "not a number",
					Path:    "x",
					Stack:   nil,
				},
			},
		},
		Test{
			Title: "index out of range",
			Input: Input{
				Object: SliceObject([]Object{
					DummyObject{1},
				}),
				Path:  "1",
				Paths: nil,
			},
			Expect: Expect{
				Object: nil,
				Err: &PathError{
					Message: "index out of range",
					Path:    "1",
					Stack:   nil,
				},
			},
		},
		Test{
			Title: "path error nested",
			Input: Input{
				Object: SliceObject([]Object{
					SliceObject([]Object{
						SliceObject([]Object{
							DummyObject{1},
						}),
					}),
				}),
				Path:  "0",
				Paths: []string{"0", "1"},
			},
			Expect: Expect{
				Object: nil,
				Err: &PathError{
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

			obj, err := testCase.Input.Object.Get(testCase.Input.Path, testCase.Input.Paths...)

			assert.Equal(testCase.Expect.Object, obj)
			assert.Equal(testCase.Expect.Err, err)
		})
	}
}

func TestSliceObject_Set(t *testing.T) {
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
				Object: SliceObject([]Object{
					DummyObject{1},
				}),
				Path:  "0",
				Paths: nil,
				BeSet: DummyObject{2},
			},
			Expect: Expect{
				Object: SliceObject([]Object{
					DummyObject{2},
				}),
				Err: nil,
			},
		},
		Test{
			Title: "success nested",
			Input: Input{
				Object: SliceObject([]Object{
					SliceObject([]Object{
						SliceObject([]Object{
							DummyObject{1},
						}),
					}),
				}),
				Path:  "0",
				Paths: []string{"0"},
				BeSet: DummyObject{2},
			},
			Expect: Expect{
				Object: SliceObject([]Object{
					SliceObject([]Object{
						DummyObject{2},
					}),
				}),
				Err: nil,
			},
		},
		Test{
			Title: "not a number",
			Input: Input{
				Object: SliceObject([]Object{
					DummyObject{1},
				}),
				Path:  "x",
				Paths: nil,
				BeSet: DummyObject{2},
			},
			Expect: Expect{
				Object: SliceObject([]Object{
					DummyObject{1},
				}),
				Err: &PathError{
					Message: "not a number",
					Path:    "x",
					Stack:   nil,
				},
			},
		},
		Test{
			Title: "index out of range",
			Input: Input{
				Object: SliceObject([]Object{
					DummyObject{1},
				}),
				Path:  "1",
				Paths: nil,
				BeSet: DummyObject{2},
			},
			Expect: Expect{
				Object: SliceObject([]Object{
					DummyObject{1},
				}),
				Err: &PathError{
					Message: "index out of range",
					Path:    "1",
					Stack:   nil,
				},
			},
		},
		Test{
			Title: "path error nested",
			Input: Input{
				Object: SliceObject([]Object{
					SliceObject([]Object{
						SliceObject([]Object{
							DummyObject{1},
						}),
					}),
				}),
				Path:  "0",
				Paths: []string{"0", "1"},
				BeSet: DummyObject{2},
			},
			Expect: Expect{
				Object: SliceObject([]Object{
					SliceObject([]Object{
						SliceObject([]Object{
							DummyObject{1},
						}),
					}),
				}),
				Err: &PathError{
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

			obj := testCase.Input.Object
			err := obj.Set(testCase.Input.BeSet, testCase.Input.Path, testCase.Input.Paths...)

			assert.Equal(testCase.Expect.Object, obj)
			assert.Equal(testCase.Expect.Err, err)
		})
	}
}

func TestSliceObject_Unwrap(t *testing.T) {
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
				Object: SliceObject([]Object{
					DummyObject{1},
				}),
			},
			Expect: Expect{
				Object: []interface{}{1},
			},
		},
		Test{
			Title: "success nested",
			Input: Input{
				Object: SliceObject([]Object{
					SliceObject([]Object{
						SliceObject([]Object{
							DummyObject{1},
						}),
					}),
				}),
			},
			Expect: Expect{
				Object: []interface{}{
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

			assert.Equal(testCase.Expect.Object, testCase.Input.Object.Unwrap())
		})
	}
}
