package main

import (
	"fmt"
	"net/http"
)

type messageHandler struct {
	message string
}

func (m *messageHandler) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(wr, m.message)
}

func main() {
	mux := http.NewServeMux()

	m1 := &messageHandler{"I am message 1"}
	mux.Handle("/message1", m1)

	m2 := &messageHandler{"I am message 2"}
	mux.Handle("/message2", m2)

	http.ListenAndServe(":8080", mux)
}
