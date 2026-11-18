package internal

// pageView returns a bounded view of the recipient queue within
// the specified [start:end] range.
//
// The returned slice shares the backing array with the original queue.
// Its capacity is restricted to its length using a full slice expression,
// preventing append operations from overwriting elements outside the view.
//
// It returns an error if the range contains negative indices,
// exceeds the queue length, or has an invalid start and end order.
func pageView(queue Queue, start, end int) (Queue, error) {
	pageLength := len(queue)

	if start < 0 || end < 0 {
		return queue, ErrIndexCannotBeNegative
	}

	if start > pageLength || end > pageLength {
		return queue, ErrIndexOverflow
	}

	if start > end {
		return queue, ErrIndexInvalid
	}

	v := queue[start:end:end]

	return v, nil
}