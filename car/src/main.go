package main

import "ssibrahimbas.com/persistent-mongo/car/src/app"

func main() {
	app.New().Init().Serve()
}
