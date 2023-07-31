package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/vaishnavi/api/model"
	// "C:\Users\vaishnavi.n.CLICK2CLOUD\Desktop\GO_Lang\File_API\model"
)

func Create_File(w http.ResponseWriter, r *http.Request) {

	file := model.File{}
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

	file := model.File{}
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
