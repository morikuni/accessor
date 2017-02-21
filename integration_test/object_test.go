package integration

import (
	"bytes"
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v2"

	"github.com/BurntSushi/toml"
	"github.com/morikuni/accessor"
	"github.com/stretchr/testify/assert"
)

func convertObject(assert *assert.Assertions, data interface{}) accessor.Accessor {
	pathString := "/friends"
	path, err := accessor.ParsePath(pathString)
	assert.Nil(err)

	acc, err := accessor.NewAccessor(data)
	assert.Nil(err)

	friend, err := acc.Get(path)
	assert.Nil(err)

	err = friend.Foreach(func(path accessor.Path, v interface{}) error {
		return friend.Set(path, v.(string)+"2")
	})
	assert.Nil(err)

	return acc
}

func TestObject_JSON(t *testing.T) {
	assert := assert.New(t)

	input := `
{
	"name": "me",
	"age": 18,
	"friends": [
		{"name": "hello"},
		{"name": "world"}
	],
	"nickname": null
}
	`
	var inputObject interface{}
	err := json.Unmarshal([]byte(input), &inputObject)
	assert.Nil(err)

	expect := `
{
	"name": "me",
	"age": 18,
	"friends": [
		{"name": "hello2"},
		{"name": "world2"}
	],
	"nickname": null
}
	`
	var exceptObject interface{}
	err = json.Unmarshal([]byte(expect), &exceptObject)
	assert.Nil(err)

	acc := convertObject(assert, inputObject)

	bs, err := json.Marshal(acc)
	assert.Nil(err)

	var result interface{}
	err = json.Unmarshal(bs, &result)
	assert.Nil(err)
	assert.Equal(exceptObject, result)
}

func TestObject_YAML(t *testing.T) {
	assert := assert.New(t)

	input := `name: me
age: 18
friends:
  - name: hello
  - name: world
nickname: null`
	var inputObject interface{}
	err := yaml.Unmarshal([]byte(input), &inputObject)
	assert.Nil(err)

	expect := `name: me
age: 18
friends:
  - name: hello2
  - name: world2
nickname: null`
	var exceptObject interface{}
	err = yaml.Unmarshal([]byte(expect), &exceptObject)
	assert.Nil(err)

	acc := convertObject(assert, inputObject)

	bs, err := json.Marshal(acc)
	assert.Nil(err)

	var result interface{}
	err = yaml.Unmarshal(bs, &result)
	assert.Nil(err)
	assert.Equal(exceptObject, result)
}

func TestObject_TOML(t *testing.T) {
	assert := assert.New(t)

	input := `
name = "me"
age = 18
time = 1992-06-18T12:34:56Z
[[friends]]
name = "hello"
[[friends]]
name = "world"
`
	var inputObject interface{}
	err := toml.Unmarshal([]byte(input), &inputObject)
	assert.Nil(err)

	expect := `
name = "me"
age = 18
time = 1992-06-18T12:34:56Z
[[friends]]
name = "hello2"
[[friends]]
name = "world2"
`
	var exceptObject interface{}
	err = toml.Unmarshal([]byte(expect), &exceptObject)
	assert.Nil(err)

	acc := convertObject(assert, inputObject)

	buf := &bytes.Buffer{}
	err = toml.NewEncoder(buf).Encode(acc.Unwrap())
	assert.Nil(err)

	var result interface{}
	err = toml.Unmarshal(buf.Bytes(), &result)
	assert.Nil(err)
	assert.Equal(exceptObject, result)
}
