package basics

/*
- Like most programming language, Go has +, -, *, / and % as operators.
- In many languages (example: JS) you can perform arithmetic operations on different types of variables,
but in Go this gives an error. (see (1))
- Shorthand Assignments & Increment and Decrement is the same as most programming language.
*/


// (1)

var x int = 42

// this line produces an error
value := float32(2.0) * x // invalid operation: mismatched types float32 and int

// you must convert int type to float32 before performing arithmetic operation
value := float32(2.0) * float32(x)