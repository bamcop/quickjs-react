package kit

func Try(err error) {
	if err != nil {
		panic(err)
	}
}
