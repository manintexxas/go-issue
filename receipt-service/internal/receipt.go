package internal

import "errors"

func Build(invoiceID ID, status Status) (string, error){
	if !validate(invoiceID, status) {
		return "", errors.New("ID or status is invalid")
	}

	return render(invoiceID, status), nil
}
