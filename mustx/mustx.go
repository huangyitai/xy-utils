package mustx

// NoError ...
func NoError(err error) {
	if err != nil {
		panic(err)
	}
}
