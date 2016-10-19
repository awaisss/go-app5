Skip to content
This repository
Search
Pull requests
Issues
Gist
 @awaisss
 Unwatch 1
  Star 0
 Fork 0 awaisss/golang-app2
 Code  Issues 0  Pull requests 0  Projects 0  Wiki  Pulse  Graphs  Settings
Branch: master Find file Copy pathgolang-app2/application.go
f0f30d4  3 hours ago
@awaisss awaisss Updates
1 contributor
RawBlameHistory     
70 lines (59 sloc)  2.02 KB

package main
 
import (
    "encoding/json"
    "log"
    "net/http"
 
    "github.com/gorilla/mux"
)
 
type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}
 
type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}
 
var people []Person
 
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Person{})
}
 
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(people)
}
 
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var person Person
    _ = json.NewDecoder(req.Body).Decode(&person)
    person.ID = params["id"]
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}
 
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(people)
}
 
func main() {
    router := mux.NewRouter()
    people = append(people, Person{ID: "1", Firstname: "awais", Lastname: "ilyas", Address: &Address{City: "Islamabad", State: "Islamabad"}})
    people = append(people, Person{ID: "2", Firstname: "ali", Lastname: "hassan"})
    router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
    router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":3000", router))
}
Contact GitHub API Training Shop Blog About
© 2016 GitHub, Inc. Terms Privacy Security Status Help