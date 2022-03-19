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

```go
array.In([]myIntType{1,2,3,4}, 2)
array.In([]float32{1.1,2,3.2,4.2}, 2)
...
```

- **Replace**: replaces an object wit another one at the 1st occurance

```go
array.Replace([]string{"a", "b", "c"}, "c", "d")
array.Replace([]complex64{1, 2, 3}, 2, 4)
...
```

- **Map**: executes a user defined function over each value in an array

```go
array.Map([]rune{'a', 'b', 'c'}, func(v rune) rune { return v + 'd' })
array.Map([]myint16{1, 2, 3}, func(v uintptr) uintptr { return v * 2 })
...
```

- **Filter**: returns a filtered array from the input array based on user defined function

```go
array.Filter([]int16{1, 2, 3}, func(v int16) bool { return v > 1 })
array.Filter([]rune{'a', 'b', 'c'}, func(v rune) bool { return v == 'b' })
...
```

- **Reduce**: returns a reduced value based on the given user defined function

```go
array.Reduce([]rune{'a', 'b', 'c'}, func(i rune, j rune) rune { return i + j })
array.Reduce([]uintptr{1, 2, 3}, func(i uintptr, j uintptr) uintptr { return i + j })
...
```

- **RemoveElem** removes an element from array

```go
array.RemoveElem([]float32{1, 2, 3}, 2)
array.RemoveElem([]byte{'a', 'b', 'c'}, 'c')
...
```

- **RemoveAt**: removes an object at given index

```go
array.RemoveAt([]int16{1, 2, 3}, 1)
array.RemoveAt([]string{"a", "b", "c"}, 1)
...
```

- **InsertAt**: inserts an object at the given index by shifting other values to right

```go
array.InsertAt([]string{"a", "b", "c"}, 2, "d")
array.InsertAt([]uintptr{1, 2, 3}, 2, 4)
...
```

- **AppendAt**: inserts a slice at given index by shifting other values to right.

```go
array.AppendAt([]rune{'a', 'b', 'c'}, 1, []rune{'d', 'e'})
array.AppendAt([]int16{1, 2, 3}, 1, []int16{3, 4})
...
```

*Code.Share.Prosper*
