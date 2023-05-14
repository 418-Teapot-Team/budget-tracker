package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type St struct {
	Name string `json:"name"`
}

func testApi(w http.ResponseWriter, r *http.Request) {

	st := St{Name: "Roman S"}

	b, _ := json.Marshal(st)

	w.Write(b)
	return

}

func main() {
	http.HandleFunc("/test", testApi)
	log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
}
