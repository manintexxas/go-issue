package internal

func appendRecipients(queue Queue, incoming Incoming) Queue {
	return append(queue, incoming...)
}