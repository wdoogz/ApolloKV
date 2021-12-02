package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var Store map[string]interface{}

// Api function
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HomePage")
}

func writeKV(w http.ResponseWriter, r *http.Request) {
	//var data string
	//var store map[string]interface{}
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal("error body not readable")
		}
		//data = string(body)
		json.Unmarshal([]byte(body), &Store)
	}
}

func kvStore(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		for k, v := range Store {
			fmt.Fprintf(w, "%s -> %s \n", k, v)
		}
	}
}

func HandleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/kv", kvStore)
	http.HandleFunc("/write", writeKV)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
