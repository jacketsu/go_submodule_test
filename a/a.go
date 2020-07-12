package main

import (
	"fmt"

	"github.com/go-modules-by-example/submodules/b"
)

const Name = b.Name

func main() {
	fmt.Print("success")
	fmt.Println(Name)
}
