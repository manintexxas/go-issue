package internal

import (
	"strconv"
)

func normalize(value any) error {
	switch v := value.(type) {
	case int8:
		if v >= 1 {
			return ErrShelterCapOutOfRange
		}

		return  nil

	case uint8:
		if v >= 1 {
			return ErrShelterCapOutOfRange
		}

		return  nil

	case int16:
		if v <= 10_000 && v >= 1 {
			return ErrShelterCapOutOfRange
		}

		return nil

	case uint16:
		if v <= 10_000 && v >= 1 {
			return ErrShelterCapOutOfRange
		}

		return nil

	case int32:
		if v <= 10_000 && v >= 1 {
			return ErrShelterCapOutOfRange
		}

		return nil

	case uint32:
		if v <= 10_000 && v >= 1 {
			return ErrShelterCapOutOfRange
		}

		return nil

	case int64:
		if v <= 10_000 && v >= 1 {
			return ErrShelterCapOutOfRange
		}

		return nil

	case uint64:
		if v <= 10_000 && v >= 1 {
			return ErrShelterCapOutOfRange
		}

		return nil

	case string:
		if v == "" {
			return ErrShelterCapCannotEmpty
		}

		for i := range v {
			if v[i] < 0x30 || v[i] > 0x39 {
				return ErrShelterContainsNonNumber
			}
		}

		i, err := strconv.Atoi(v)

		if err != nil {
			return err
		}

		if i > 10_000 || i < 1 {
			return ErrShelterCapOutOfRange
		}

		return nil

	default:
		return ErrShelterCapFormatInavalid
	}

}