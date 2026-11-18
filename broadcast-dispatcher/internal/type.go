package internal

import "errors"

type (
	Channel		[4]bool
	Queue		[]string
	Incoming	[]string
	SnapShot	[]string
)

var (
	ErrChannelIndexOutOfRange		= errors.New("channel index must be between 0 - 3")
	ErrIndexCannotBeNegative		= errors.New("index cannot be negative number")
	ErrIndexOverflow				= errors.New("index cannot be high than page")
	ErrIndexInvalid 				= errors.New("start index cannot over end index")
)