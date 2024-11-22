# Range

The range form of the for loop iterates over a slice, array, or map.

When ranging over an array or slice, two values are returned for each iteration. 
The first is the index, and the second is a copy of the element at that index.

When ranging over a map, two values are returned for each iteration.
The first is the key, and the second is a copy of the value.

You can skip the index, key, or value by assigning to _.

```aiignore
list := make([]int, 10)
for i, _ := range list {
    // do something with the index i
}

for _, value := range list {
    // do something with a copy of the value
}
```

If you only want the index or key, you can omit the second variable.

`for i := range list`