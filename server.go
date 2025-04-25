package main

import (
	"net/http"
)

const port = "8080"

func serverInit() error {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	mux.Handle("/", http.FileServer(http.Dir(".")))

	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
