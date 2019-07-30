package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// type msg string

// func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
// 	fmt.Fprint(resp, m)
// }

// type httpHandler struct {
// 	message string
// }

// func (h httpHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
// 	fmt.Fprint(resp, h.message)
// }

func productsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("Product %s", id)
	fmt.Fprint(w, response)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id := vars["id"]
	response := fmt.Sprintf("This is a main page")
	fmt.Fprint(w, response)
}

func main() {
	// portNumber := "8181"
	// hostPath := "localhost:" + portNumber
	// // msgHandler := msg("Hello from Web Server in Go")

	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "hello.html")
	// })

	// http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "About Page")
	// })
	// http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Contact Page")
	// })
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Index Page")
	// })

	// h1 := httpHandler{message: "Test Handler"}
	// h2 := httpHandler{message: "Test Handler 2"}
	// http.Handle("/testh1", h1)
	// http.Handle("/testh2", h2)

	// fmt.Println("---> Server is listening on port: ", hostPath)
	// // http.ListenAndServe(hostPath, msgHandler)
	// http.ListenAndServe(hostPath, nil)

	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "static/hello.html")
	// })
	// http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "static/test.html")
	// })
	// fmt.Println("Server is listening...")
	// http.ListenAndServe(":8181", nil)

	router := mux.NewRouter()
	router.HandleFunc("/products/{id:[0-9]+}", productsHandler)
	router.HandleFunc("/", mainHandler)

	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
