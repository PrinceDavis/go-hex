package main

import (
	"fmt"

	"github.com/PrinceDavis/hex/internal/adapters/core/arithmetic"
)

func main() {
	arithAdapter := arithmetic.NewAdapter()
	sum, err := arithAdapter.Addition(1, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(sum)
}
