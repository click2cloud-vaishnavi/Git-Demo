package router

import (
	"github.com/gorilla/mux"
	"github.com/vaishnavi/api/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/create", controller.Create_File).Methods("POST")
	router.HandleFunc("/get/{name}", controller.Get_File).Methods("GET")
	router.HandleFunc("/update", controller.Update_File).Methods("PUT")
	router.HandleFunc("/delete/{name}", controller.Delete_File).Methods("DELETE")
	// http.ListenAndServe(":5000", router)
	return router

}
