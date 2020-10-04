package main

import (
	app "github.com/Pranjalya/nuxt-go-neo4j-starter/server/app"
)

func main() {
	myApp := app.Init()
	myApp.InitRoutes()
	myApp.Run()
}
