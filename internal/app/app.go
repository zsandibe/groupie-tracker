package app

import (
	"errors"
	"log"
	d "zsandibe/internal/delivery"
	s "zsandibe/internal/server"
)

func Start() error {
	ourHandler := d.NewHandler()
	server := s.Server(*ourHandler)
	log.Println("Server is running...\nhttp://localhost:8080/")
	err := server.ListenAndServe()
	if err != nil {
		return errors.New(err.Error())
	}
	return err
}
