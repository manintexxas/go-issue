package internal

// Config updates the enabled state of a broadcast channel.
//
// The channel index must be between 0 and 3.
// It returns the updated channel configuration or an error
// if the channel index is out of range.
func Config(current Channel, channelIndex int, enabled bool) (Channel, error) {
	return setChannelEnabled(current, channelIndex, enabled)
}

// Recipients appends incoming recipients to the existing queue.
//
// It preserves the order of incoming recipients and returns
// the updated queue. The returned slice must be used because
// appending may allocate a new backing array.
func Recipients(queue Queue, incoming Incoming) Queue {
	return appendRecipients(queue, incoming)
}

// View returns a bounded view of the recipient queue within
// the specified [start:end] range.
//
// The returned slice shares the backing array with the original queue.
// Its capacity is restricted to its length to prevent append operations
// from overwriting elements outside the selected range.
//
// It returns an error if the specified range is invalid.
func View(queue Queue, start, end int) (Queue, error) {
	return pageView(queue, start, end)
}

// Audit returns an independent snapshot of the recipient queue
// within the specified [start:end] range.
//
// The snapshot uses a separate backing array, so modifications
// to the returned slice do not affect the original queue.
//
// It returns an error if the specified range is invalid.
func Audit(queue Queue, start, end int) (SnapShot, error) {
	return snapShot(queue, start, end)
}