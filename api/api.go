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
	fmt.Fprintln(w, "ApolloKV")
	jsonOut, _ := json.MarshalIndent(Store, "", "  ")
	fmt.Fprintln(w, string(jsonOut))
	log.Printf("[INFO] %s request from %s", r.Method, r.RemoteAddr)
}

func writeKV(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Body is not readable\n %s", r.Body)
			log.Printf("[ERROR] body not readable, %s", r.Body)
		}
		//data = string(body)
		err = json.Unmarshal([]byte(body), &Store)
		if err != nil {
			fmt.Fprintf(w, "ERROR: Cannot send data due to error (%s).", err)
			log.Printf("[WARN] %s request threw client side error (%s)", r.Method, err)
		} else {
			log.Printf("[INFO] %s request from %s successful", r.Method, r.RemoteAddr)
		}
	}
	if r.Method == "DELETE" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Body is not readable\n %s", r.Body)
			log.Printf("[ERROR] body not readable, %s", r.Body)
		}
		if Store[(string(body))] == nil {
			fmt.Fprintf(w, "Please pick a valid Key to delete!\n")
		} else {
			delete(Store, string(body))
			log.Printf("[INFO] %s request from %s successful", r.Method, r.RemoteAddr)
			log.Printf("[WARN] %s KV pair was deleted", string(body))
			fmt.Fprintf(w, "KV pair %s was deleted successfully\n", string(body))
		}
	}
}

func kvStore(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		for k, v := range Store {
			fmt.Fprintf(w, "%s -> %s\n", k, v)
		}
	}
}

func HandleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/kv", kvStore)
	http.HandleFunc("/write", writeKV)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
