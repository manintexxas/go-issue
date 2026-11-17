package internal

import "time"

// CreateRecord processes a source, validates its shelter ID and capacity,
// rejects duplicate IDs, and stores a new PENDING record in Dict.
func CreateRecord(c CapacitySource) (*ShelterRecord, error) {
	tg, rawID, rawCap, err := ProccessSource(c)
	if err != nil {
		return nil, err
	}

	id, err := validateShelterID(rawID)
	if err != nil {
		return nil, err
	}

	capacity, err := validateShelterCap(rawCap)
	if err != nil {
		return nil, err
	}

	if _, exists := Dict[id]; exists {
		return nil, ErrIDAlreadyReported
	}

	rec := &ShelterRecord{
		Update: time.Now(),
		Tag:    tg,
		ID:     id,
		Cap:    capacity,
		Status: PENDING,
	}

	Dict[id] = rec
	return rec, nil
}