
package internal

import "errors"

// ID represents an invoice identifier.
type ID string

// Status represents the current state of an invoice.
type Status uint8

const (
	// PAID indicates that the invoice has been paid.
	PAID Status = 1 << iota

	// PENDING indicates that the invoice is awaiting payment.
	PENDING

	// FAILED indicates that the invoice payment has failed.
	FAILED

	// UNKNOWN represents an unspecified invoice status.
	UNKNOWN
)

var (
	// ErrIDLengthNotValid indicates that the invoice ID
	// does not have the required length.
	ErrIDLengthNotValid = errors.New("invoice ID must be exactly 8 bytes")

	// ErrIDPrefixNotValid indicates that the invoice ID
	// does not have the required prefix.
	ErrIDPrefixNotValid = errors.New("ID prefix is invalid")

	// ErrIDContainsNonNumber indicates that the invoice ID
	// sequence contains non-digit characters.
	ErrIDContainsNonNumber = errors.New("ID sequence number must contain digits only")

	// ErrStatusNotValid indicates that the provided invoice
	// status is not supported.
	ErrStatusNotValid = errors.New("Status not valid")
)
