package challenges

// on error print the error and exit
func check(err error) {
	if err != nil {
		panic(err)
	}
}
