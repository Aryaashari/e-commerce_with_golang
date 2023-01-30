package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type Config struct {
	appName, appEnv, appPort string
}

func (server *Server) Initialize(config Config) {
	fmt.Println(config.appName)
	server.Router = mux.NewRouter()
	server.InitRouter()
}

func (server *Server) Run(config Config) {
	fmt.Println("Running", config.appEnv, "Server On Port", config.appPort)
	log.Fatal(http.ListenAndServe(config.appPort, server.Router))
}

func getEnv(key, defaultName string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultName
}

func Run() {
	server := Server{}
	config := Config{}

	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	config.appName = getEnv("APP_NAME", "E-Commerce")
	config.appEnv = getEnv("APP_ENV", "dev")
	config.appPort = getEnv("APP_PORT", ":2023")

	server.Initialize(config)
	server.Run(config)
}
