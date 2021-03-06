package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type QnrResponse struct {
    Id uuid.UUID `json:"id"`
    Name string `json:"name"`
    Body string `json:"body"`
    Email string `json:"email"`
}

var QnrResponses []QnrResponse

func deleteQnrResponse(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := uuid.Parse(vars["id"])

    if (err != nil) {
        log.Fatal(err)
        return
    }

    for index, response := range QnrResponses {
        if response.Id == id {
            QnrResponses = append(QnrResponses[:index], QnrResponses[index+1:]...)
        }
    }
}

func addQnrResponse(w http.ResponseWriter, r *http.Request) {
    // get the body of the POST request
    // return the response body as a string
    reqBody, _ := ioutil.ReadAll(r.Body)
    var response QnrResponse
    json.Unmarshal(reqBody, &response)

    response.Id, _ = uuid.NewRandom()

    // append response to global variable
    QnrResponses = append(QnrResponses, response)

    json.NewEncoder(w).Encode(response)
}

func getQnrResponse(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key, err := uuid.Parse(vars["id"])

    if (err != nil) {
        log.Fatal(err)
        return
    }

    // Loop through each response, and if the ID matches, return the response
    for _, response := range QnrResponses {
        if response.Id == key {
            json.NewEncoder(w).Encode(response)
        }
    }
}

func getAllQnrResponses(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(QnrResponses)
}

func handleRequests() {
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/api/v1/responses", getAllQnrResponses).Methods("GET")

    router.HandleFunc("/api/v1/responses", addQnrResponse).Methods("POST")
    router.HandleFunc("/api/v1/responses/{id}", getQnrResponse).Methods("GET")

    router.HandleFunc("/api/v1/responses/{id}", deleteQnrResponse).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
    QnrResponses = []QnrResponse{}

    handleRequests()
}
