package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/joho/godotenv"
	"github.com/vikashvverma/upworkdemo/app"
	"github.com/vikashvverma/upworkdemo/routes/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}

	app.Init()
	muxRouter := router.Router()
	n := negroni.New()
	n.UseHandler(muxRouter)

	log.Printf("Server listening on http://localhost:%s/", port)
	n.Run(fmt.Sprintf(":%s", port))
}
