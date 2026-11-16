package internal

import "errors"

type (
	ID     string
	Status uint8
)

const (
	PAID Status = 1 << iota
	PENDING
	FAILED
	UNKNOWN
)

var (
	ErrIDLengthNotValid		= errors.New("ID length mush 8 char")
	ErrIDPrefixNotValid		= errors.New("ID prefix is invalid")
	ErrIDContainsNonNumber	= errors.New("ID sequence number must contain digits only")

	ErrStatusNotValid		= errors.New("Status not valid")
)


