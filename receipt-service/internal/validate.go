package internal

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

func validateStatus(status Status) error {
	switch status {
		case PAID, PENDING, FAILED:
			return nil

		default:
			return ErrStatusNotValid
	}
}