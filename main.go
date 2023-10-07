package main

import (
	"fmt"
	"go-basic/routes"
	"net/http"
)

func main() {
	mainRoutes := routes.Routes()
	http.Handle("/", mainRoutes)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
