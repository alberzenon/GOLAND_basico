package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "¡Hola bienvenidos a mi pagina web !")
}

func main() {
	var h Hello
	http.ListenAndServe("localhost:5000", h)
}

