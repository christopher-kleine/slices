# slices - Chain modifiers on Slices

## Table of Content

- About
- Installation
- Usage
- Examples

### About

`slices` let's you use modifiers on slices in a rather simple way. It uses the
"Options Pattern". It also provides functions to perfom checks or transformations
to the slice such as slicing and grouping.

### Installation

Using `go modules`: `go get github.com/christopher-kleine/slices`

### Usage

The packages is divided into seperate packages based on their type. Those are:

- `str` for `string`
- `sint` for `signed int (int)`
- `sint8` for `signed int8 (int8)`
- `sint16` for `signed int16 (int16)`
- `sint32` for `signed int32 (int32)`
- `sint64` for `signed int64 (int64)`
- `fl32` for `float32`
- `fl64` for `float64`
- `usint` for `unsigned int (uint)` 
- `usint8` for `unsigned int (uint8)` 
- `usint16` for `unsigned int (uint16)` 
- `usint32` for `unsigned int (uint32)` 
- `usint64` for `unsigned int (uint64)` 

Simply import them like this:

```go
import "github.com/christopher-kleine/slices/str"
```

### Examples

1. Filtering log file for `"[WARNING]"` and `[ERROR]`.

```go
lines := []string{
	"[WARNING] user entered wrong password",
	"[INFO] system successfully started",
	"[DEBUG] cpu-load is at 40%",
	"[ERROR] user interrupted update",
}

filtered := str.Chain(lines,
	str.Select(str.Contains("[WARNING]", "[ERROR]")),
)
// Result:
// 	"[WARNING] user entered wrong password",
//	"[ERROR] user interrupted update"

```

2. Take only unique Tags recieved from an HTTP-Request

```go
receivedTags := []string{"foo", "bar", "Foo", "world"}

uniqueTags := str.Chain(receivedTags,
	str.ToLower,
	str.Uniques,
)
// Result:
// "foo", "bar", "world"
```
