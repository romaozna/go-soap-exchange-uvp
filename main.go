package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"main/src/controller"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func getPort() int {
	port := 7540
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	envPort := os.Getenv("TODO_PORT")

	if len(envPort) > 0 {
		if pport, err := strconv.ParseInt(envPort, 10, 32); err == nil {
			port = int(pport)
		}
	}

	return port
}

func main() {
	r := chi.NewRouter()
	r.Get("/", controller.GetStatus)
	r.Get("/api/v0/archive/{deviceId}/{pipeId}", controller.GetArchive)
	r.Get("/api/v0/devices/", controller.GetDevices)
	r.Get("/api/v0/devices/{deviceId}/", controller.GetDeviceById)

	serverPort := getPort()
	log.Println(fmt.Sprintf("Адрес сервера: %d", serverPort))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serverPort), r))
}
