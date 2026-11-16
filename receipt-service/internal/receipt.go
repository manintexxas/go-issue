
package internal

// Build creates a formatted invoice receipt.
//
// It validates the invoice ID and status before generating
// the receipt.
//
// If either validation fails, Build returns an empty string
// and the corresponding validation error.
//
// If both validations succeed, it returns the generated
// receipt and nil error.
func Build(invoiceID ID, status Status) (string, error) {
	if err := validateID(invoiceID); err != nil {
		return "", err
	}

	if err := validateStatus(status); err != nil {
		return "", err
	}

	return render(invoiceID, status), nil
}
