package main

import "SimpleWebProject/internal/app"
import _ "SimpleWebProject/docs"

// @title Users Example App
// @version 1.0
// @description Simple User Service

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	app.RunApp()
}
