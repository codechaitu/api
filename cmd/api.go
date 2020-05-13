package main

import "github.com/rideonstyle/api/server"

func main(){
	router := server.CreateRouter()
	server.StartServer(router)
}
