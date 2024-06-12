package main

import (
	"github.com/IIGabriel/training-sheet-crud/internal/server"
	"log"
)

func main() {
	s := server.NewServer()
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
