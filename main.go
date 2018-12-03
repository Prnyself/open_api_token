package main

import (
	"fmt"
	"log"
	"net/http"
	"open_api_token/router"
	"open_api_token/settings"
)

func main() {
	controller := router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
		Handler:        controller,
		ReadTimeout:    settings.ReadTimeout,
		WriteTimeout:   settings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
