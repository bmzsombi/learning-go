package richterscale

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// describeEarthquake returns the "description" of a given magnitude value on the Richter scale.
func describeEarthquake(magnitude float32) string {
	// INSERT YOUR CODE HERE
	var strength string
	switch {
	case magnitude < 2.0:
		strength = "micro"
	case magnitude >= 2.0 && magnitude < 3.0:
		strength = "very minor"
	case magnitude >= 3.0 && magnitude < 4.0:
		strength = "minor"
	case magnitude >= 4.0 && magnitude < 5.0:
		strength = "light"
	case magnitude >= 5.0 && magnitude < 6.0:
		strength = "moderate"
	case magnitude >= 6.0 && magnitude < 7.0:
		strength = "strong"
	case magnitude >= 7.0 && magnitude < 8.0:
		strength = "major"
	case magnitude >= 8.0 && magnitude < 10.0:
		strength = "great"
	case magnitude >= 10.0:
		strength = "massive"
	}
	return strength
}
