package main

import (
	"banco-de-dados/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{ID}", servidor.BuscarUsuario).Methods(http.MethodGet)

	fmt.Println("escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
