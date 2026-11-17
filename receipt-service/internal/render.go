package internal

import "fmt"

// render formats an invoice identifier and its current status into a
// human-readable summary string.
func render(invoiceID ID, status Status) string {
	label := "UNKNOWN"
	switch status {
	case PAID:
		label = "PAID"
	case PENDING:
		label = "PENDING"
	case FAILED:
		label = "FAILED"
	}

	return fmt.Sprintf("Invoice: %s\nStatus: %s", invoiceID, label)
}