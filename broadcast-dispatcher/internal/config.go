package internal

// setChannelEnabled updates the enabled state of a channel at the specified index.
//
// It returns the updated channel configuration without modifying the
// original array. An error is returned if the channel index is outside
// the valid range of 0 to 3.
func setChannelEnabled(current Channel, channelIndex int, enabled bool) (Channel, error) {
	if channelIndex > 3 || channelIndex < 0 {
		return current, ErrChannelIndexOutOfRange
	}

	current[channelIndex] = enabled

	return current, nil
}