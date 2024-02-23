package main

import (
	"fmt"
	t "time"
)

func main() {
	fmt.Printf("%s | Lets fresh ur Golang skills!", t.Now().Format("02/01/2006 15:04:05"))
}
