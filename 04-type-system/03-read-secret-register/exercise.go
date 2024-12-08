package readsecretregister

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// parseChannelControlRegister constructs 4 octets (8-bit long uint) based on the parameter register.
func parseChannelControlRegister(charCtrl uint32) (uint8, uint8, uint8, uint8) {
	// INSERT YOUR CODE HERE
	byte1 := uint8(charCtrl & 0xFF)
	byte2 := uint8((charCtrl >> 8) & 0xFF)
	byte3 := uint8((charCtrl >> 16) & 0xFF)
	byte4 := uint8((charCtrl >> 24) & 0xFF)
	return byte2, byte3, byte4, byte1
}
