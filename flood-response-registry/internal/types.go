package internal

import "errors"

type (
	second			= uint8

	ShelterID       string
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
	Creat	second
	Update	second

	ID		ShelterID
	Cap		ShelterCapacity
	Status	ReportStatus
}

type CapacitySource interface {
	SourceName() 	string
	RawShelterID() 	any
	RawCapacity() 	any
}

var (
	ErrIDContainsNonNumber			= errors.New("ID sequence number must contain digits only")
	ErrIDLengthNotValid				= errors.New("Shelter ID must be exactly 9 bytes")
	ErrIDFormatInvalid				= errors.New("Shelter ID format is invalid")

	ErrShelterCapFormatInavalid		= errors.New("Shelter cap format is invalid")
	ErrShelterCapOutOfRange 		= errors.New("Shelter capacity must be between 1 and 10000")
	ErrShelterCapCannotEmpty		= errors.New("Shelter capacity cannot be empty")
	ErrShelterContainsNonNumber		= errors.New("Shelter cap must contain digits only")
)