# go-func-collections

commonly used functions to reduce code bulk

[![Go](https://github.com/ric-v/go-func-collections/actions/workflows/go.yml/badge.svg)](https://github.com/ric-v/go-func-collections/actions/workflows/go.yml)
[![CodeQL](https://github.com/ric-v/go-func-collections/actions/workflows/codeql.yml/badge.svg)](https://github.com/ric-v/go-func-collections/actions/workflows/codeql.yml)
[![CodeFactor](https://www.codefactor.io/repository/github/ric-v/go-func-collections/badge)](https://www.codefactor.io/repository/github/ric-v/go-func-collections)
[![Go Report Card](https://goreportcard.com/badge/github.com/ric-v/go-func-collections)](https://goreportcard.com/report/github.com/ric-v/go-func-collections)
[![GoDoc](https://godoc.org/github.com/narqo/go-badge?status.svg)](https://pkg.go.dev/github.com/ric-v/go-func-collections/array)

---

## Usage

Import the package to the project

```bash
go get github.com/ric-v/go-func-collections
```

```go
package somepackage

import (
    "fmt"
    "github.com/ric-v/go-func-collections/array"
)

func somefunc(args ...any) {
    /*
    ... doing some important stuff ...
    */

    names := []string{"James", "John", "Jim", "Jack"}

    // check if a name is in the list
    exists, index, name := array.In(names, "Jack") 
    fmt.Println(exists, index, name) // output: true 3 Jack

    // same method can be applied for any data types and 
    // even data types that are replacements for a primitive data type
    type myType int

    array.In([]myType{1,2,3,4}, 2) // output: true 1 2
    array.In([]int{1,2,3,4}, 2) // output: true 1 2
    array.In([]uintptr{1,2,3,4}, 2) // output: true 1 2
    array.In([]float32{1.1,2.1,3.1,4.1}, 2.1) // output: true 1 2.1
    array.In([]float32{1.1,2.1,3.1,4.1}, 2.2) // output: false 0 2.2

    array.Replace([]int16{1, 2, 3}, 2, 4) // [1 4 3]
    array.RemoveAt([]string{"a", "b", "c"}, "c") // [a b]
    array.Map([]int16{1, 2, 3}, func(v int16) int16 { return v * 2 }) // [2 4 6]
    array.Filter([]byte{'a', 'b', 'c'}, func(v byte) bool { return v == 'b' }) // [b]
    array.Reduce(
        []complex64{1, 2, 3}, 
        func(i complex64, j complex64) complex64 { return i + j }
    ) // (6+0i)
    array.InsertAt([]uintptr{1, 2, 3}, 2, 4) // [1 2 4 3]
    array.AppendAt([]string{"a", "b", "c"}, 1, []string{"d", "e"}) // [a d e b c]
}
```

## Available functions

- **In**: validates if given object is available in the slice
- **Replace**: replaces an object wit another one at the 1st occurance
- **Map**: executes a user defined function over each value in an array
- **Filter**: returns a filtered array from the input array based on user defined function
- **Reduce**: returns a reduced value based on the given user defined function
- **RemoveAt**: removes an object by finding the index
- **InsertAt**: inserts an object at the given index by shifting other values to right
- **AppendAt**: inserts a slice at given index by shifting other values to right.

*Code.Share.Prosper*
