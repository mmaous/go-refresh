## Data types in Go

1. **Numeric Types:**
- `int`: Signed integers (int8, int16, int32, int64).
- `uint`: Unsigned integers (uint8, uint16, uint32, uint64).
- `float32`: 32-bit floating-point numbers.
- `float64`: 64-bit floating-point numbers.
- `complex64`, `complex128`: Complex numbers.

2. **String Type:**
- `string`: Represents a sequence of characters.

3. **Boolean Type:**
- `bool`: Represents boolean values (`true` or `false`).

4. **Derived Types:**
- `array`: Fixed-size collection of elements of the same type.
- `slice`: Dynamic-size, flexible view into elements of an array.
- `map`: Unordered collection of key-value pairs.
- `struct`: Composite data type that groups together variables under a single name.
- `pointer`: Represents the memory address of a variable.

5. **Special Types:**
- `chan`: Channel type used for communication between goroutines.
- `func`: Function type.
- `interface`: Defines a set of method signatures.
- `interface{}`: The empty interface, representing any type.
- `nil`: Represents the absence of a value or a pointer that does not point to any memory address.

6. **Predeclared Constants:**
- `true`, `false`: Boolean constants.
- `iota`: Constant generator for incrementing values.
- `nil`: Represents the zero value for pointers, interfaces, channels, and function types.

7. **Predeclared Variables:**
- `int`, `float64`, `string`, etc.: Predefined variables for certain types.
