package internal

func pageView(queue Queue, start, end int) (Queue, error) {
	pageLength	:= len(queue)

	if start < 0 || end < 0 {
		return queue, ErrIndexCannotBeNegative
	}

	if start > pageLength ||end > pageLength {
		return queue, ErrIndexOverflow
	}

	if start > end {
		return queue, ErrIndexInvalid
	}

	v := queue[start:end:end]

	return v, nil
}