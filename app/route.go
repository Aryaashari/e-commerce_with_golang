package app

import "github.com/Aryaashari/e-commerce_with_golang/app/controller"

func (server *Server) InitRouter() {
	server.Router.HandleFunc("/", controller.Home).Methods("GET")
}
