# Primitive Types

Go's basic types are:

```
bool // not convertable to/from integers!!

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32 - represents a Unicode code point

float32 float64

complex64 complex128
```
The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.

### Type Conversion

assignment between items of different type requires an explicit conversion.
```aiignore
var i int = 24
var j int = 24
	
var f float64 = float64(i + j)  // good

var f float64 = i + j  // cannot use i + j (value of type int) as float64 value in variable declaration
```