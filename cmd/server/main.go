package main

import "github.com/zeimedee/saber/internal/router"

func main() {
	port := ":4000"

	router := router.SetupRouter()

	router.Run(port)
}
