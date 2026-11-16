
package internal

// validateID validates an invoice ID.
//
// The invoice ID must:
//   - Have exactly 8 bytes.
//   - Start with the "INV-" prefix.
//   - Contain exactly 4 ASCII digits after the prefix.
//
// It returns an error when the invoice ID does not satisfy
// the required format. Otherwise, it returns nil.
func validateID(invoiceID ID) error {
	if len(invoiceID) != 8 {
		return ErrIDLengthNotValid
	}

	if invoiceID[:4] != "INV-" {
		return ErrIDPrefixNotValid
	}

	for i := 4; i < 8; i++ {
		if invoiceID[i] < 0x30 || invoiceID[i] > 0x39 {
			return ErrIDContainsNonNumber
		}
	}

	return nil
}

// validateStatus validates an invoice status.
//
// Only PAID, PENDING, and FAILED are accepted.
// It returns ErrStatusNotValid for any other status.
func validateStatus(status Status) error {
	switch status {
	case PAID, PENDING, FAILED:
		return nil

	default:
		return ErrStatusNotValid
	}
}
