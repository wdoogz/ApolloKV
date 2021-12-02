package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Api function
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
	fmt.Println(r.Method)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("error body not readable")
	}
	fmt.Println(string(body))
}

func HandleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
