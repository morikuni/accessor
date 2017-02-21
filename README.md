# accessor
[![Build Status](https://travis-ci.org/morikuni/accessor.svg?branch=master)](https://travis-ci.org/morikuni/accessor)
[![GoDoc](https://godoc.org/github.com/morikuni/accessor?status.svg)](https://godoc.org/github.com/morikuni/accessor)
[![Go Report Card](https://goreportcard.com/badge/github.com/morikuni/accessor)](https://goreportcard.com/report/github.com/morikuni/accessor)

accessor provides accessor for map/slice objects like json, yaml, toml etc.

## Install

```
go get github.com/morikuni/accessor
```

## Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/morikuni/accessor"
)

func main() {
	text := `
	{
		"name": "morikuni",
		"friends": [
			{ "name": "hello" },
			{ "name": "world" }
		]
	}`

	var obj interface{}
	err := json.Unmarshal([]byte(text), &obj)
	if err != nil {
		log.Fatal(err)
	}

	acc, err := accessor.NewAccessor(obj)
	if err != nil {
		log.Fatal(err)
	}

	path, err := accessor.ParsePath("/friends")
	if err != nil {
		log.Fatal(err)
	}

	friends, err := acc.Get(path)
	if err != nil {
		log.Fatal(err)
	}

	err = friends.Foreach(func(path accessor.Path, value interface{}) error {
		return friends.Set(path, value.(string)+"_accessor")
	})
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(os.Stdout).Encode(acc.Unwrap())
	if err != nil {
		log.Fatal(err)
	}
}
```

will output

```json
$ go run main.go | jq .
{
  "friends": [
    {
      "name": "hello_accessor"
    },
    {
      "name": "world_accessor"
    }
  ],
  "name": "morikuni"
}

```
