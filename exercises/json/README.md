# JSON

Go offers built-in support for JSON encoding and decoding, including to and from built-in and custom data types.

Package [encoding/json](https://pkg.go.dev/encoding/json) implements encoding and decoding of JSON as defined in RFC 7159. 
The mapping between JSON and Go values is described in the documentation for the `Marshal` and `Unmarshal` functions.

## Encoding

To encode JSON data we use the `Marshal` function.

`func Marshal(v interface{}) ([]byte, error)`

Given the Go data structure, `Message`,

```aiignore
type Message struct {
    Name string
    Body string
    Time int64
}
```

and an instance of `Message`

`m := Message{"Alice", "Hello", 1294706395881547000}`

we can marshal a JSON-encoded version of `m` using `json.Marshal`:

`b, err := json.Marshal(m)`

If all is well, `err` will be `nil` and `b` will be a `[]byte` containing this JSON data:

```
b == []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
```

Only data structures that can be represented as valid JSON will be encoded:

* JSON objects only support strings as keys; to encode a Go map type it must be of the form `map[string]T` (where `T` is any Go type supported by the json package).

* Channel, complex, and function types cannot be encoded.

* Cyclic data structures are not supported; they will cause `Marshal` to go into an infinite loop.

* Pointers will be encoded as the values they point to (or `null` if the pointer is `nil`).

The json package only accesses the exported fields of struct types (those that begin with an uppercase letter). 
Therefore, only the exported fields of a struct will be present in the JSON output.

## Decoding

To decode JSON data we use the `Unmarshal` function.

`func Unmarshal(data []byte, v interface{}) error`

We must first create a place where the decoded data will be stored

`var m Message`

and call `json.Unmarshal`, passing it a `[]byte` of JSON data and a pointer to `m`

`err := json.Unmarshal(b, &m)`

If `b` contains valid JSON that fits in `m`, after the call `err` will be `nil` and the data from `b` will have been 
stored in the struct `m`, as if by an assignment like:

```
m = Message{
        Name: "Alice",
        Body: "Hello",
        Time: 1294706395881547000,
    }
```

How does `Unmarshal` identify the fields in which to store the decoded data? For a given JSON key "Foo", 
`Unmarshal` will look through the destination struct’s fields to find (in order of preference):

* An exported field with a tag of "Foo" (see the Go spec for more on struct tags),

* An exported field named "Foo", or

* An exported field named "FOO" or "FoO" or some other case-insensitive match of "Foo".

What happens when the structure of the JSON data doesn’t exactly match the Go type?

```
b := []byte(`{"Name":"Bob","Food":"Pickle"}`)
var m Message
err := json.Unmarshal(b, &m)
```

`Unmarshal` will decode only the fields that it can find in the destination type. In this case, only the `Name` 
field of `m` will be populated, and the `Food` field will be ignored. This behavior is particularly useful when you 
wish to pick only a few specific fields out of a large JSON blob. It also means that any unexported fields in 
the destination struct will be unaffected by `Unmarshal`.

## Streaming Encoders and Decoders

The json package provides `Decoder` and `Encoder` types to support the common operation of reading and writing 
streams of JSON data. The `NewDecoder` and `NewEncoder` functions wrap the `io.Reader` and `io.Writer` interface types.

```
func NewDecoder(r io.Reader) *Decoder
func NewEncoder(w io.Writer) *Encoder
```

Here’s an example program that reads a series of JSON objects from standard input, removes all but the `Name` field 
from each object, and then writes the objects to standard output:

```
package main

import (
    "encoding/json"
    "log"
    "os"
)

func main() {
    dec := json.NewDecoder(os.Stdin)
    enc := json.NewEncoder(os.Stdout)
    for {
        var v map[string]interface{}
        if err := dec.Decode(&v); err != nil {
            log.Println(err)
            return
        }
        for k := range v {
            if k != "Name" {
                delete(v, k)
            }
        }
        if err := enc.Encode(&v); err != nil {
            log.Println(err)
        }
    }
}
```

Due to the ubiquity of `Readers` and `Writers`, these `Encoder` and `Decoder` 
types can be used in a broad range of scenarios, such as reading and writing to HTTP connections, WebSockets, or files.

## JSON Tags

JSON Tags can be used when marshaling or unmarshaling.
As an example, consider the following example:

```aiignore
package main

import (
    "fmt"
    "encoding/json"
)

type User struct {
    Name string `json:"full_name"`
    Age int `json:"age,omitempty"`
    Active bool `json:"-"`
    lastLoginAt string
}

func main() {
        u, err := json.Marshal(User{Name: "John", Age: 25, Active: true, lastLoginAt: "today"})
        if err != nil {
            panic(err)
        }
        fmt.Println(string(u)) // {"full_name":"John"}
}
```

In this example, you can see that:

* You can specify a JSON tag json:"full_name" and specify a custom name for your JSON field
* When specifying json:"<field name>,omitempty" as a json tag, the field will be discarded in the JSON output if the 
value has a zero value. This is mostly used to discard optional JSON fields in the output
* When specifying json:"-" as a json tag, the field will be removed altogether from the JSON output. 
This is used when you want to keep the field in your Go struct but not in your JSON output