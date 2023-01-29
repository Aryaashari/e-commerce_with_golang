package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize() {
	fmt.Println("E-Commerce")
	server.Router = mux.NewRouter()
	server.InitRouter()
}

func (server *Server) Run(addr string) {
	fmt.Println("Running Server On Port", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func Run() {
	server := Server{}
	server.Initialize()
	server.Run(":2023")
}
