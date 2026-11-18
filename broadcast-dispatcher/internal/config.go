package internal

func setChannelEnabled(current Channel, channelIndex int, enabled bool) (Channel, error) {
	if channelIndex > 3 || channelIndex < 0 {
		return current, ErrChannelIndexOutOfRange
	}

	current[channelIndex] = enabled

	return current, nil
}