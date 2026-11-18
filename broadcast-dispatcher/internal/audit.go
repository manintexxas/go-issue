package internal

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