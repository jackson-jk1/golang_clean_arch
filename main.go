package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()
	fmt.Println(config.ConectionString)
	fmt.Println("rodando api")

	r := router.ReturnRoutes()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
