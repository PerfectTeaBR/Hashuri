package main

import (
	"fmt"
	"net/http"
)

func hashuriPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hashuri")
}

func main() {
	// create a server
	http.HandleFunc("/", hashuriPage)
	fmt.Println("Servidor rodando na porta 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro:", err)
	}
}
