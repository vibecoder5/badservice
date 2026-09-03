package main

import (
	"styleguide/layering/app"
)

func main() {
	searchAPI := app.New()

	searchAPI.Start()
}
