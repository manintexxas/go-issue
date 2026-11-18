package internal

// appendRecipients appends incoming recipients to the existing queue.
//
// It preserves the order of incoming recipients and returns the updated queue.
// The append operation may reuse the existing backing array or allocate
// a new one if the queue's capacity is insufficient.
func appendRecipients(queue Queue, incoming Incoming) Queue {
	return append(queue, incoming...)
}