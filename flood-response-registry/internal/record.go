package internal

import "time"

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

	cap, err := validateShelterCap(cp)

	if err != nil {
		return nil, err
	}

	s := &ShelterRecord{
		Update: time.Now(),
		Tag:	tg,
		ID:		id,
		Cap: 	cap,
		Status: PENDING,
	}

	Dict[id] = s

	return s, nil
}
