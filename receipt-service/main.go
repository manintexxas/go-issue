package main

import (
	"fmt"
	"math/rand"
	in "receipt-service/internal"
)

func IDmaker(prefix string) func() in.ID {
	return func () in.ID {
		min := 1000
		max := 9999

		gen := rand.Intn(max-min+1) + min

		return in.ID(fmt.Sprintf("%s-%d", prefix, gen))
	}
}

func main() {
	var status in.Status

	func() {
		num := [4]in.Status{1, 2, 4, 8}

		n	:= num[rand.Intn(len(num))]

		status = in.Status(n)
	}()

	invID := IDmaker("INV")

	fmt.Print(in.Build(invID(), status))
}

