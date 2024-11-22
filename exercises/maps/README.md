# Maps

A map maps keys to values. It is unordered.

The zero value of a map is nil. A nil map has no keys, nor can keys be added. 
If a key is attempted to be retrieved the zero value for value type will be returned.
```aiignore
var m map[string]int

n := m["One"] // returns 0
m["One"] = 1  // panic - nil map
```

Map keys must be comparable (no slices, maps, or funcs)

## Create a Map

### Make

The make function returns a map of the given type, initialized and ready for use.

```aiignore
m := make(map[string]int)
m["test"] = 1

n := make(map[string]int, 100) // creates a map that can hold at least 100 elements
```

### Map Literal

You can also create a map with a map literal.

```aiignore
m := map[string]int{"test": 1}
n := map[string]int{} // empty map
```

If the top-level type is just a type name, you can omit it from the elements of the literal.

```aiignore
type User struct {
    Name string
    Age  int
}

m := map[int]User{
    1: {"John", 25},
    2: {"Jane", 25},
}
```

## Mutating Maps

Insert or update an element in map `m`:

`m[key] = elem`

Retrieve an element:

`elem = m[key]`

**If key is not in the map, then elem is the zero value for the map's element type.**

Delete an element:

`delete(m, key)`

Test that a key is present with a two-value assignment:

`elem, ok = m[key]`

If key is in `m`, `ok` is `true`. If not, `ok` is `false`. 
(`ok` is the convention in Golang, but it could be anything)

**If key is not in the map, then elem is the zero value for the map's element type.**

Note: If elem or ok have not yet been declared you could use a short declaration form:

`elem, ok := m[key]`

### Comparable

Maps are NOT comparable.

```aiignore
m1 := map[int]string{1: "One"}
m2 := map[int]string{1: "One"}

fmt.Printf("%#v", m1 == m2) // invalid operation: m1 == m2 (map can only be compared to nil)
```

### Passed by

Maps can be treated as if passed by reference.



