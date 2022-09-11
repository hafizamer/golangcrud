package main

import (
	"net/http"

	"golangCrud/controllers/patientController"
)

func main() {
	http.HandleFunc("/", patientController.Index)
	http.HandleFunc("/patient", patientController.Index)
	http.HandleFunc("/patient/index", patientController.Index)
	http.HandleFunc("/patient/add", patientController.Add)
	//http.HandleFunc("/patient/edit", patientController.Edit)
	//http.HandleFunc("/patient/delete", patientController.Delete)

	http.ListenAndServe(":3000", nil)
}
