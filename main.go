package main

var appImages DisplayedImages

func main() {
	build, err := Build{}.build()

	if err != nil {
		println(err.Error())
		return
	}

	Run(build)
}
