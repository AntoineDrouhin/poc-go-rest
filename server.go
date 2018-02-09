package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Index)
	router.HandleFunc("/help", Index)

	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Print("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		`API usage Guide : 
			/help
			/helloWorld
			/todos
			/todos/{todoId}
		`)
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Todo Index !")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintf(w, "Todo ID : %q ! ", todoId)
}
