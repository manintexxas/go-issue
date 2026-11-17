package internal

func (m MobileRequest) SourceName() Tags {
	return MOBILE
}

func (m MobileRequest) RawID() any {
	return m.ShelterID
}

func (m MobileRequest) RawCap() any {
	return  m.Capacity
}

func (l LegacyCSVRow) SourceName() Tags {
	return CSV
}

func (l LegacyCSVRow) RawID() any {
	return l.ShelterID
}

func (l LegacyCSVRow) RawCap() any {
	return  l.Capacity
}

func (p PartnerPayload) SourceName() Tags {
	return NGO
}

func (p PartnerPayload) RawID() any {
	return p.ShelterID
}

func (p PartnerPayload) RawCap() any {
	return  p.Capacity
}

func ProccessSource(c CapacitySource) (Tags, any, any, error) {
	if c.RawID() == nil || c.RawCap() == nil{
		return 0, nil, nil, ErrRequestEmpty
	}

	return c.SourceName(), c.RawID(), c.RawCap(), nil
}