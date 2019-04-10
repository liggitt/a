package main

import (
	"fmt"
	"github.com/liggitt/a/helpers"
	"github.com/liggitt/b"
	"github.com/liggitt/c"
	"github.com/liggitt/d"
)

func main() {
	b.Cake()
	fmt.Println(c.CakeOrDeath())
	fmt.Println(d.Ingredients())
	helpers.PrintSuccess()
}
