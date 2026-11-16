package internal

import ("fmt"
		"errors"
)

func Build(invoiceID ID, status Status) (string, error){
	if !validate(invoiceID, status) {
		return "", errors.New("ID or status cannot empty")
	}

	return fmt.Sprintf("%s", render(invoiceID, status)), nil
}
