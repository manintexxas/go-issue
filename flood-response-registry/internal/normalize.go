package internal

import "strconv"

// normalize accepts the numeric types used by the supported intake sources,
// converts them to uint16, and enforces the inclusive capacity range 1..10000.
// String values must contain decimal digits only and must also fit in int.
func normalize(value any) (uint16, error) {
	switch v := value.(type) {
	case int8:
		if v < 1 {
			return 0, ErrShelterCapOutOfRange
		}

		return uint16(v), nil

	case uint8:
		if v < 1 {
			return 0, ErrShelterCapOutOfRange
		}

		return uint16(v), nil

	case int16:
		if !(v <= 10_000 && v >= 1) {
			return 0, ErrShelterCapOutOfRange
		}

		return uint16(v), nil

	case uint16:
		if !(v <= 10_000 && v >= 1) {
			return 0, ErrShelterCapOutOfRange
		}

		return v, nil

	case int32:
		if !(v <= 10_000 && v >= 1) {
			return 0, ErrShelterCapOutOfRange
		}

		return uint16(v), nil

	case uint32:
		if !(v <= 10_000 && v >= 1) {
			return 0, ErrShelterCapOutOfRange
		}

		return uint16(v), nil

	case int:
		if !(v <= 10_000 && v >= 1) {
			return 0, ErrShelterCapOutOfRange
		}

		return uint16(v), nil

	case uint:
		if !(v <= 10_000 && v >= 1) {
			return 0, ErrShelterCapOutOfRange
		}

		return uint16(v), nil

	case int64:
		if !(v <= 10_000 && v >= 1) {
			return 0, ErrShelterCapOutOfRange
		}

		return uint16(v), nil

	case uint64:
		if !(v <= 10_000 && v >= 1) {
			return 0, ErrShelterCapOutOfRange
		}

		return uint16(v), nil

	case string:
		if v == "" {
			return 0, ErrShelterCapCannotEmpty
		}

		for i := range v {
			if v[i] < 0x30 || v[i] > 0x39 {
				return 0, ErrShelterContainsNonNumber
			}
		}

		i, err := strconv.Atoi(v)

		if err != nil {
			return 0, ErrIDContainsNonNumber
		}

		if i > 10_000 || i < 1 {
			return 0, ErrShelterCapOutOfRange
		}

		return uint16(i), nil

	default:
		return 0, ErrShelterCapFormatInavalid
	}
}
