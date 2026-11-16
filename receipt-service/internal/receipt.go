package internal

import "fmt"

func Build(invoiceID ID, status Status) {
	if !validate(invoiceID, status) {
		fmt.Printf("ID or status cannot empty ")
	}

	fmt.Printf("%s", render(invoiceID, status))
}
