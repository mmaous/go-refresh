package basics

import "fmt"

/**
- Go applications are organized in packages.
- A package is a collection of source files located in the same directory.
- All source files in a directory must share the same package name.
- The recommended style of naming in Go is that identifiers will be named using `camelCase`,
except for those meant to be accessible across packages which should be `PascalCase`.
- When a package is imported, only entities (functions, types, variables, constants) whose names start with a capital letter can be used / accessed.
*/

// PublicKey & SayHey are Accessible when package is imported
var PublicKey = "your-public-key"

func SayHey() {
	fmt.Println("Hey!!")
}

// privateKey & SayHey Not accessible when package is imported
var privateKey = "your-private-key"

func sayHello() {
	fmt.Println("Hello!!")
}
