package main

import (
	"fmt"
    "log"
	"net/http"
    "bytes"
    "io/ioutil"

	"github.com/gorilla/mux"
)

func createChat(w http.ResponseWriter, r *http.Request){

    vars := mux.Vars(r)
	var application_token string
    application_token = vars["application_token"]
  
    var jsonStr = []byte(`{}`)
	response, err := http.Post(
        "http://localhost:3000/api/v1/applications/"+application_token+"/chats",
         "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
        fmt.Print(err.Error())
    }
    data, _ := ioutil.ReadAll(response.Body)
    w.Write(data)
}


func main() {
    router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}