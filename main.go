package main

import "BookStore-API/app"

func main() {
	// Init server
	server := app.App{}
	server.Initialize()
	server.Run()
}
