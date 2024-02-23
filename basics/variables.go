package basics
// NOTE: this file can't be executed
/*
- Go is statically-typed, which means all variables must have a defined type at compile-time.
*/

// Variables can be defined by explicitly specifying a type:
var bathrooms int // Explicitly typed


// You can use an initializer, and the compiler will automatically assign the variable type based on the initializer.
bedrooms := 2

// U can always later on update its value
bedrooms = 3
bathrooms = 2

// keep in mind, That u can only pass the value of its initial type
bathrooms = false // throws error at compile time as its expecting 'int'

// CONSTANT

// Constants hold a piece of data just like variables, but their value cannot change during the execution of the program.
const stories = 2 // Defines a numeric constant 'stories' with the value of 2
