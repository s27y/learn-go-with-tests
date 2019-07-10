package main

import (
	"fmt"
	"net/http"
)

// type Handler interface {
// 	ServeHTTP(http.ResponseWriter, *http.Request)
// }

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	fmt.Fprint(w, GetPlayerScore(player))
	return

}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}
	if name == "Floyd" {
		return "10"
	}
	return ""
}
