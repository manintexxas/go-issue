package internal

import "fmt"

func render(invoiceID ID, status Status) string {
	s := ""

	switch status {
	case PAID:
		s = "PAID"

	case PENDING:
		s = "PENDING"

	case FAILED:
		s = "FAILED"
	}

	output := fmt.Sprintf("Invoice: %s\nStatus: %s", invoiceID, s)

	return output
}