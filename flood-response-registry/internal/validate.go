package internal

func validateShelterID(raw any) error {
	r, ok := raw.(string)

	if !ok {
		return ErrIDFormatInvalid
	}

	if len(r) != 9 {
		return ErrIDLengthNotValid
	}

	if err := normalize(r[:5]); err != nil {
		return ErrIDContainsNonNumber
	}

	return nil
}

func validateShelterCap(raw any) error {
	if err := normalize(raw); err != nil {
		return err
	}

	return nil
}