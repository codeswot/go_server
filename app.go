package main

import (
	"fmt"
	"net/http"

	genfunctions "github.com/elmubarak/go_crud/genFunctions"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := httprouter.New()
	router.GET("/", genfunctions.FetchAll)
	router.GET("/:id", genfunctions.FetchById)
	router.POST("/", genfunctions.CreateStudent)
	router.PUT("/:id", genfunctions.UpdateStudent)
	router.DELETE("/:id", genfunctions.DeleteStudent)
	fmt.Println("Server listening to port 3000")
	http.ListenAndServe(":3000", router)
}
