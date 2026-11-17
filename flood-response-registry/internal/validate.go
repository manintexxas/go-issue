package internal

// prefixCheck reports whether the supplied four-byte prefix belongs to a
// supported intake source.
func prefixCheck(prefix string) bool {
	return prefix == "MBL-" || prefix == "CSV-" || prefix == "NGO-"
}


// validateShelterID checks the raw ID type, required byte length, source
// prefix, and four-digit sequence. It returns the normalized ShelterID when
// every check succeeds.
func validateShelterID(raw any) (ShelterID, error) {
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
	for i := 4; i < 8; i++ {
		if r[i] < 0x30 || r[i] > 0x39 {
			return "", ErrIDContainsNonNumber
		}
	}
	return ShelterID(r), nil
}

// validateShelterCap delegates raw capacity conversion and range checking to
// normalize, then converts the result to the ShelterCapacity domain type.
func validateShelterCap(raw any) (ShelterCapacity, error) {
	r, err := normalize(raw)
	if err != nil {
		return 0, err
	}
	return ShelterCapacity(r), nil
}
