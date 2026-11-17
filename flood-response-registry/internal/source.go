package internal

// SourceName returns the tag identifying a mobile request.
func (m MobileRequest) SourceName() Tags {
	return MOBILE
}

// RawID returns the unvalidated shelter ID from a mobile request.
func (m MobileRequest) RawID() any {
	return m.ShelterID
}

// RawCap returns the unvalidated capacity from a mobile request.
func (m MobileRequest) RawCap() any {
	return m.Capacity
}

// SourceName returns the tag identifying a legacy CSV row.
func (l LegacyCSVRow) SourceName() Tags {
	return CSV
}

// RawID returns the unvalidated shelter ID from a legacy CSV row.
func (l LegacyCSVRow) RawID() any {
	return l.ShelterID
}

// RawCap returns the unvalidated capacity from a legacy CSV row.
func (l LegacyCSVRow) RawCap() any {
	return l.Capacity
}

// SourceName returns the tag identifying an NGO partner payload.
func (p PartnerPayload) SourceName() Tags {
	return NGO
}

// RawID returns the unvalidated shelter ID from an NGO partner payload.
func (p PartnerPayload) RawID() any {
	return p.ShelterID
}

// RawCap returns the unvalidated capacity from an NGO partner payload.
func (p PartnerPayload) RawCap() any {
	return p.Capacity
}

// ProccessSource extracts the source tag, raw shelter ID, and raw capacity.
// It rejects a nil ID or capacity before the values reach the validators.
func ProccessSource(c CapacitySource) (Tags, any, any, error) {
	id := c.RawID()
	capacity := c.RawCap()

	if id == nil || capacity == nil {
		return 0, nil, nil, ErrRequestEmpty
	}

	return c.SourceName(), id, capacity, nil
}
