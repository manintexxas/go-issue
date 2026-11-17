// Package internal implements shelter intake processing, validation,
// record creation, and the in-memory shelter registry.
package internal

import (
	"errors"
	"time"
)

// second is an alias for time.Time used by ShelterRecord.Update.
type (
	second = time.Time

	// ShelterID represents a shelter's source-prefixed identifier.
	ShelterID string
	// ShelterCapacity represents a validated shelter capacity.
	ShelterCapacity uint16
	// ReportStatus represents the review state of a shelter record.
	ReportStatus uint8
	// Tags identifies the source from which a shelter record was received.
	Tags uint8
)

// Dict is the in-memory registry of shelter records, indexed by shelter ID.
// Records remain available only for the lifetime of the running application.
var Dict = make(map[ShelterID]*ShelterRecord)

// Report status values used by shelter records. Each value occupies a distinct
// bit, although the application uses them as individual states.
const (
	UNKNOWN ReportStatus = 1 << iota
	PENDING
	VERIFIED
	REJECTED
)

// Source tags identify the intake channel associated with a shelter record.
const (
	MOBILE Tags = 1 << iota
	CSV
	NGO
)

// ShelterRecord stores the normalized information accepted by the intake
// process, along with its source, current status, and last update time.
type ShelterRecord struct {
	Update second

	Tag    Tags
	ID     ShelterID
	Cap    ShelterCapacity
	Status ReportStatus
}

// MobileRequest represents intake data received through the mobile source.
// Its capacity is represented as an int64 before normalization.
type MobileRequest struct {
	ShelterID string
	Capacity  int64
}

// LegacyCSVRow represents a row imported from the legacy CSV source.
// Its capacity is represented as a string before normalization.
type LegacyCSVRow struct {
	ShelterID string
	Capacity  string
}

// PartnerPayload represents intake data received from an NGO partner.
// Its any-typed fields allow the intake layer to validate incoming value types.
type PartnerPayload struct {
	ShelterID any
	Capacity  any
}

// CapacitySource describes the methods required to extract source identity,
// shelter ID, and raw capacity from an intake source.
type CapacitySource interface {
	SourceName() Tags
	RawID() any
	RawCap() any
}

// Intake and validation errors returned by the internal package.
var (
	ErrIDContainsNonNumber = errors.New("ID sequence number must contain digits only")
	ErrIDLengthNotValid    = errors.New("Shelter ID must be exactly 8 bytes")
	ErrIDFormatInvalid     = errors.New("Shelter ID format is invalid")
	ErrIDPrefixInvalid     = errors.New("Shelter ID prefix is invalid")

	ErrShelterCapFormatInavalid = errors.New("Shelter cap format is invalid")
	ErrShelterCapOutOfRange     = errors.New("Shelter capacity must be between 1 and 10000")
	ErrShelterCapCannotEmpty    = errors.New("Shelter capacity cannot be empty")
	ErrShelterContainsNonNumber = errors.New("Shelter cap must contain digits only")

	ErrIDAlreadyBeenReport = errors.New("ID already have been use")
	ErrRequestEmpty        = errors.New("ID or cap request cannot be empty")
)
