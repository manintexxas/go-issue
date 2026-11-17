package internal

import "time"

// CreatedRecord processes a source, validates its shelter ID and capacity,
// rejects duplicate IDs, and stores a new PENDING record in Dict. On success,
// it returns the same record pointer stored in the registry.
func CreatedRecord(c CapacitySource) (*ShelterRecord, error) {
	tg, i, cp, err := ProccessSource(c)

	if err != nil {
		return nil, err
	}

	id, err := validateShelterID(i)

	if err != nil {
		return nil, err
	}

	if _, ok := Dict[id]; ok {
		return nil, ErrIDAlreadyBeenReport
	}

	capacity, err := validateShelterCap(cp)

	if err != nil {
		return nil, err
	}

	s := &ShelterRecord{
		Update: time.Now(),
		Tag:    tg,
		ID:     id,
		Cap:    capacity,
		Status: PENDING,
	}

	Dict[id] = s

	return s, nil
}
