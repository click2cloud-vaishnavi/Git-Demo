package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type File struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func main() {
	fmt.Println("Welcome to the File")
	r := mux.NewRouter()

	r.HandleFunc("/create", Create_File).Methods("POST")
	r.HandleFunc("/get/{name}", Get_File).Methods("GET")
	r.HandleFunc("/update", Update_File).Methods("PUT")
	r.HandleFunc("/delete/{name}", Delete_File).Methods("DELETE")

	http.ListenAndServe(":5000", r)

}

func Create_File(w http.ResponseWriter, r *http.Request) {

	file := File{}
	_ = json.NewDecoder(r.Body).Decode(&file)
	name := file.Name
	data := file.Data

	_ = ioutil.WriteFile(name, []byte(data), 0644)

	fmt.Fprintf(w, "File created successfully...!!!!!")
}

func Get_File(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/get/"):]

	data, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(data))

}

func Update_File(w http.ResponseWriter, r *http.Request) {

	file := File{}
	_ = json.NewDecoder(r.Body).Decode(&file)

	name := file.Name
	data := file.Data
	_ = ioutil.WriteFile(name, []byte(data), 0644)

	fmt.Fprintf(w, "File updated successfully")

}

func Delete_File(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Path[len("/delete/"):]
	_ = os.Remove(name)

	fmt.Fprintf(w, "File deleted successfully")
}
