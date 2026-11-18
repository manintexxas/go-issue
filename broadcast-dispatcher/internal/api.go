package internal 

func Config(current Channel, channelIndex int, enabled bool) (Channel, error) {
	return setChannelEnabled(current, channelIndex, enabled)
}

func Recipients(queue Queue, incoming Incoming) Queue {
	return appendRecipients(queue, incoming)
}

func View(queue Queue, start, end int) (Queue, error) {
	return pageView(queue, start, end)
}

func Audit(queue Queue, start, end int) (SnapShot, error) {
	return snapShot(queue, start, end)
}