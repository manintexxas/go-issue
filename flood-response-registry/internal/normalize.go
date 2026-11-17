package internal

import "strconv"

const (
	minShelterCap = 1
	maxShelterCap = 10_000
)

// fromSigned validates that a signed capacity value falls within the inclusive
// range minShelterCap..maxShelterCap and converts it to uint16.
func fromSigned(v int64) (uint16, error) {
	if v < minShelterCap || v > maxShelterCap {
		return 0, ErrShelterCapOutOfRange
	}
	return uint16(v), nil
}

// fromUnsigned validates that an unsigned capacity value falls within the
// inclusive range minShelterCap..maxShelterCap and converts it to uint16.
func fromUnsigned(v uint64) (uint16, error) {
	if v < minShelterCap || v > maxShelterCap {
		return 0, ErrShelterCapOutOfRange
	}
	return uint16(v), nil
}

// normalize accepts the numeric types used by the supported intake sources,
// converts them to uint16, and enforces the inclusive capacity range 1..10000.
// String values must contain decimal digits only and must also fit in int.
func normalize(value any) (uint16, error) {
	switch v := value.(type) {
	case int8:
		return fromSigned(int64(v))

	case int16:
		return fromSigned(int64(v))

	case int32:
		return fromSigned(int64(v))

	case int:
		return fromSigned(int64(v))

	case int64:
		return fromSigned(v)

	case uint8:
		return fromUnsigned(uint64(v))

	case uint16:
		return fromUnsigned(uint64(v))

	case uint32:
		return fromUnsigned(uint64(v))

	case uint:
		return fromUnsigned(uint64(v))

	case uint64:
		return fromUnsigned(v)

	case string:
		if v == "" {
			return 0, ErrShelterCapCannotEmpty
		}
		for i := 0; i < len(v); i++ {
			if v[i] < 0x30 || v[i] > 0x39 {
				return 0, ErrShelterContainsNonNumber
			}
		}
		n, err := strconv.Atoi(v)
		if err != nil {
			return 0, ErrShelterCapOutOfRange
		}
		return fromSigned(int64(n))

	default:
		return 0, ErrShelterCapFormatInvalid
	}
}