package internal

// snapShot creates an independent copy of the specified queue range.
//
// It validates the requested range against the queue length and returns
// an error if the indices are invalid. The returned snapshot uses a
// separate backing array, so modifications to the snapshot do not affect
// the original queue.
func snapShot(queue Queue, start, end int) (SnapShot, error) {
	pageLength := len(queue)

	if start < 0 || end < 0 {
		return SnapShot{}, ErrIndexCannotBeNegative
	}

	if start > pageLength || end > pageLength {
		return SnapShot{}, ErrIndexOverflow
	}

	if start > end {
		return SnapShot{}, ErrIndexInvalid
	}

	queue = queue[start:end]

	destination := make(SnapShot, len(queue))

	copy(destination, queue)

	return destination, nil
}