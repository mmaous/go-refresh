package main

import (
	"fmt"
	b "github.com/mmaous/go-refresh/basics"
	t "time"
)

func main() {
	fmt.Printf("%s | Lets fresh ur Golang skills!\n", t.Now().Format("02/01/2006 15:04:05"))
	b.SayHey()
	fmt.Println("Your Public key is", b.PublicKey)
	// output
	/**
	> go run .\main.go
	23/02/2024 11:15:28 | Lets fresh ur Golang skills!
	Hey!!
	Your Public key is your-public-key
	*/

	// b.sayHello()
	// fmt.Println("Your Public key is", b.privateKey)
	// output (crashes)
	/**
	> go run .\main.go
	# command-line-arguments
	.\main.go:15:4: undefined: b.sayHello
	.\main.go:16:38: undefined: b.privateKey
	*/

	// --------------------

	RandomCar := b.AudiQ8
	fmt.Printf("Car{number: non-accessible, Model: %s, Color: %s, DistanceInKm: %.2f}", RandomCar.Model, RandomCar.Color, RandomCar.DistanceInKm)
	// output: Car{number: non-accessible, Model: Audi Q8, Color: Matte-black, DistanceInKm: 1205.60}
}
