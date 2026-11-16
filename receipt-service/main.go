
package main

import (
	"fmt"
	"math/rand"
	in "receipt-service/internal"
)

// IDmaker creates an invoice ID generator using the given prefix.
//
// It returns a closure that generates a new ID each time it is called.
// The generated ID consists of the provided prefix followed by
// a randomly generated four-digit number.
func IDmaker(prefix string) func() in.ID {
	return func() in.ID {
		min := 1000
		max := 9999

		gen := rand.Intn(max-min+1) + min

		return in.ID(fmt.Sprintf("%s-%d", prefix, gen))
	}
}

// main initializes an invoice status, generates an invoice ID,
// and builds the corresponding receipt.
//
// The invoice status is selected randomly from the available
// status values.
//
// If receipt generation fails, the error is printed and the
// function returns without printing a receipt.
func main() {
	var status in.Status

	func() {
		num := [4]in.Status{1, 2, 4, 8}

		n := num[rand.Intn(len(num))]

		status = in.Status(n)
	}()

	invID := IDmaker("INV")

	makeReceipt, err := in.Build(invID(), status)

	if err != nil {
		fmt.Printf("Failed to build receipt: %v\n", err)
		return
	}

	fmt.Print(makeReceipt)
}
