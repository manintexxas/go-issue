package internal

import (
	"errors"
	"time"
)

type (
	second			= time.Time

	ShelterID       string
	ShelterCapacity uint16
	ReportStatus    uint8
	Tags			uint8
)

var Dict = make(map[ShelterID]*ShelterRecord)

const (
	UNKNOWN ReportStatus = 1 << iota
	PENDING
	VERIFIED
	REJECTED
)

const (
	MOBILE Tags	= 1 << iota
	CSV
	NGO
)

type ShelterRecord struct {
	Update	second

	Tag		Tags
	ID		ShelterID
	Cap		ShelterCapacity
	Status	ReportStatus
}

type MobileRequest struct {
    ShelterID	string
    Capacity	int64
}

type LegacyCSVRow struct {
    ShelterID 	string
    Capacity  	string
}

type PartnerPayload struct {
    ShelterID	any
    Capacity	any
}

type CapacitySource interface {
	SourceName()	Tags
	RawID()			any
	RawCap()		any
}

var (
	ErrIDContainsNonNumber			= errors.New("ID sequence number must contain digits only")
	ErrIDLengthNotValid				= errors.New("Shelter ID must be exactly 8 bytes")
	ErrIDFormatInvalid				= errors.New("Shelter ID format is invalid")
	ErrIDPrefixInvalid				= errors.New("Shelter ID prefix is invalid")

	ErrShelterCapFormatInavalid		= errors.New("Shelter cap format is invalid")
	ErrShelterCapOutOfRange 		= errors.New("Shelter capacity must be between 1 and 10000")
	ErrShelterCapCannotEmpty		= errors.New("Shelter capacity cannot be empty")
	ErrShelterContainsNonNumber		= errors.New("Shelter cap must contain digits only")

	ErrIDAlreadyBeenReport			= errors.New("ID already have been use")

	ErrRequestEmpty					= errors.New("ID or cap request cannot be empty")

)