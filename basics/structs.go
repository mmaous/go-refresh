package basics

import "math/rand"

/* What are structs in Go?
In Go, a struct is a sequence of named elements called fields, each field has a name and type.
- The name of a field must be unique within the struct.
- Structs can be compared with classes in the Object-Oriented Programming paradigm.
- You create a new struct by using the `type` and `struct` keywords, and explicitly define the name and type of the fields.
- You can create an instance of a struct without using the field names, as long as you define the fields **IN ORDER**.
- Fields that don't have an initial value assigned, will have their zero value.
*/

// Car | Define a Struct
type Car struct {
	number       int
	Model        string
	Color        string
	DistanceInKm float32
}

// AudiQ8 | To create a new instance of Car define the fields using their field name in any order:
var AudiQ8 = Car{
	number:       12253,
	Model:        "Audi Q8",
	Color:        "Matte-black",
	DistanceInKm: 1205.6,
}

// "New" functions : They might remind you of constructors in other languages, but in Go they are just regular functions.
/*
Using New functions can have the following advantages:
- Validation of the given values.
- Handling of default-values.
- Since New functions are often declared in the same package of the structs they initialize,
they can initialize even private fields of the struct
*/

func NewCar(model string) Car {
	return Car{
		number:       rand.Int(),
		Model:        model,
		Color:        "white",
		DistanceInKm: 0,
	}
}

// Use the ``new`` builtin to create new instances of structs as a last resort.
