// Package internal provides data types and errors for the
// regional emergency broadcast dispatcher.
package internal

import "errors"

type (
	// Channel represents the enabled state of four broadcast channels.
	Channel [4]bool

	// Queue represents a collection of recipient identifiers.
	Queue []string

	// Incoming represents a collection of recipient identifiers
	// to be added to an existing queue.
	Incoming []string

	// SnapShot represents a snapshot of a selected recipient queue range.
	SnapShot []string
)

var (
	// ErrChannelIndexOutOfRange indicates that the channel index
	// is outside the valid range of 0 to 3.
	ErrChannelIndexOutOfRange = errors.New("channel index must be between 0 - 3")

	// ErrIndexCannotBeNegative indicates that a range index is negative.
	ErrIndexCannotBeNegative = errors.New("index cannot be negative number")

	// ErrIndexOverflow indicates that a range index exceeds the queue length.
	ErrIndexOverflow = errors.New("index cannot be high than page")

	// ErrIndexInvalid indicates that the start index is greater than the end index.
	ErrIndexInvalid = errors.New("start index cannot over end index")
)