package patientController

import (
	"fmt"
	"golangCrud/entities"
	"golangCrud/models"
	"net/http"
	"text/template"
)

var patientModel = models.NewPatientModel()

func Index(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Hello")
	temp, err := template.ParseFiles("views/patient/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)

}

func Add(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/patient/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var patient entities.Patient
		patient.Name = request.Form.Get("name")
		patient.Nick = request.Form.Get("nick")
		patient.Male_female = request.Form.Get("male_female")
		patient.Place_of_birth = request.Form.Get("place_of_birth")
		patient.Date_of_birth = request.Form.Get("date_of_birth")
		patient.Address = request.Form.Get("adress")
		patient.Phone = request.Form.Get("phone")
		fmt.Println(patient)
		patientModel.Create(patient)
		data := map[string]interface{}{
			"message": "data saved",
		}
		temp, _ := template.ParseFiles("views/patient/add.html")

		temp.Execute(response, data)
	}

}

func Edit(response http.ResponseWriter, request *http.Request) {

}

func Delete(response http.ResponseWriter, request *http.Request) {

}
