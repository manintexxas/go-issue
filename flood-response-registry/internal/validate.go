package internal

func prefixCheck(prefix string) bool{
	switch prefix {
	case "MBL-", "CSV-", "NGO-":
		return true

	default:
		return false
	}
}

func validateShelterID(raw any) (ShelterID, error)  {
	r, ok := raw.(string)

	if !ok {
		return "", ErrIDFormatInvalid
	}

	if len(r) != 8 {
		return "", ErrIDLengthNotValid
	}

	if !prefixCheck(r[:4]) {
		return "", ErrIDPrefixInvalid
	}
	
	for i := 4; i < 8; i++{
		if r[i] < 0x30 || r[i] > 0x39 {
				return "", ErrShelterContainsNonNumber
			}
	} 

	return ShelterID(r), nil
}

func validateShelterCap(raw any) (ShelterCapacity, error) {
	r, err := normalize(raw)
	
	if err != nil {
		return 0, err
	}

	return ShelterCapacity(r), nil
}

