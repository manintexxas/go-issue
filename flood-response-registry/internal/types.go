package internal

import "time"

type (
	LegacyShelterID = string

	ShelterID       LegacyShelterID
	ShelterCapacity uint16
	ReportStatus    uint8
)

const (
	UNKNOWN ReportStatus = 1 << iota
	PENDING
	VERIFIED
	REJECTED
)

type ShelterRecord struct {
	ID		ShelterID
	Cap		ShelterCapacity
	Status	ReportStatus

	Date	time.Time
}

type CapaitySource interface {
	SourceName() 	string
	RawShelterID() 	any
	RawCapacity() 	any
}