package main

import (
	"log"
	"os"

	"github.com/Ruginis-AV/finish-six.sprint/internal/handlers"
	"github.com/Ruginis-AV/finish-six.sprint/internal/server"
	"github.com/Ruginis-AV/finish-six.sprint/internal/service"
	"github.com/Ruginis-AV/finish-six.sprint/pkg/morse"
)

func main() {
	log := log.New(os.Stdout, "", 0)

	converter := morse.NewConverter(morse.DefaultMorse)

	mainService := service.New(converter)

	mainHandler := handlers.New(log, mainService)

	serv := server.New(log, "localhost", "8080", 10, 5, 15, mainHandler)

	serv.Start()
}
