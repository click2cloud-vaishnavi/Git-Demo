package main

import (
	"fmt"
	"net/http"

	router "github.com/vaishnavi/File_API/controller"
)

func main() {
	fmt.Println("Welcome to the File")
	r := router.Router()

	http.ListenAndServe(":5000", r)

}
