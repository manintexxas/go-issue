package internal

func Build(invoiceID ID, status Status) (string, error){
	if err := validateID(invoiceID); err != nil {
		return "", err
	}

	if err := validateStatus(status); err != nil {
		return "", err
	}

	return render(invoiceID, status), nil
}
