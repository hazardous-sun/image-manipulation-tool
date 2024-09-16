package main

var UnsavedProgress bool

func main() {
	build, err := Build{}.build()

	if err != nil {
		println(err.Error())
		return
	}

	UnsavedProgress = false

	Run(build)
}
