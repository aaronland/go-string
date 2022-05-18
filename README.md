# go-string

Go package providing methods for creating and manipulating common string types.

There are many string libraries. This one is mine.

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/aaronland/go-string.svg)](https://pkg.go.dev/github.com/aaronland/go-string)

## Example

```
package main

import (
	"fmt"
	"github.com/aaronland/go-string/random"
)

func main() {

	opts := random.DefaultOptions()
	opts.Length = 40
	opts.Chars = 20

	s, _ := random.String(opts)

	fmt.Println(s)
}
```

_Error handling omitted for the sake of brevity._

## Tools

### randomstr

```
./bin/randomstr -h
Usage of ./bin/randomstr:
  -alphanumeric
    	Only include alpha-numeric characters (this causes the -ascii flag to be set to true)
  -ascii
    	Only include ASCII characters
  -chars int
    	Minimum length of the random string, in characters
  -length int
    	Minimum length of the random string, in bytes (default 32)
```	
