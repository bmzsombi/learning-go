package secretprotocolheader

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// createPublishFixHeader constructs an octet (8-bit long byte) based on its three arguments and the fix QoS setting.
func createPublishFixHeader(isFirstAttempt bool, isBroadcasted bool, isSecure bool) byte {
	// INSERT YOUR CODE HERE
	QoS := byte(0b01)

	var header byte
	header |= 0b010 << 5

	if isFirstAttempt {
		header |= 1 << 4
	} else {
		header |= 0 << 4
	}

	header |= QoS << 2

	if isBroadcasted {
		header |= 1 << 1
	} else {
		header |= 0 << 1
	}

	if isSecure {
		header |= 1 << 0
	} else {
		header |= 0 << 0
	}

	return header
}
