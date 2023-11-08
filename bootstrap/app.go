package bootstrap

type Applciation struct {
}

func App() Applciation {
	app := &Applciation{}

	return *app
}
