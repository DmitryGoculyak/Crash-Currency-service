package main

import "Crash-Currency-service/internal/container"

func main() {
	app := container.Build()

	app.Run()
}
