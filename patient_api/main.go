package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for patient - file
type Patient struct {
	PatientId   string   `json:"patientid"`
	PatientName string   `json:"patientname"`
	Age         int      `json:"age"`
	Disease     *Disease `json:"disease"`
}

type Disease struct {
	Fullname string `json:"fullname"`
}

//fake DB
var patients []Patient

// middleware, helper - file
func (c *Patient) IsEmpty() bool {
	// return c.PatientId == "" && c.PatientName == ""
	return c.PatientName == ""
}

func main() {
	fmt.Println("API - LearnCodeOnline.in")
	r := mux.NewRouter()

	//seeding
	patients = append(patients, Patient{PatientId: "2", PatientName: "Radha Patil", Age: 29, Disease: &Disease{Fullname: "Radha Patil"}})
	patients = append(patients, Patient{PatientId: "4", PatientName: "Niraj Pawar", Age: 19, Disease: &Disease{Fullname: "Niraj Pawar"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/patients", getAllPatients).Methods("GET")
	r.HandleFunc("/patient/{id}", getOnePatient).Methods("GET")
	r.HandleFunc("/patient", createOnePatient).Methods("POST")
	r.HandleFunc("/patient/{id}", updateOnePatient).Methods("PUT")
	r.HandleFunc("/patient/{id}", deleteOnePatient).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

func getAllPatients(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all information")
	w.Header().Set("Content-Type", "applicatioan/json")
	json.NewEncoder(w).Encode(patients)
}

func getOnePatient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one patient info")
	w.Header().Set("Content-Type", "applicatioan/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through patients, find matching id and return the response
	for _, patient := range patients {
		if patient.PatientId == params["id"] {
			json.NewEncoder(w).Encode(patient)
			return
		}
	}
	json.NewEncoder(w).Encode("No patient found with given id")
	return
}

func createOnePatient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one patient")
	w.Header().Set("Content-Type", "applicatioan/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}

	var patient Patient
	_ = json.NewDecoder(r.Body).Decode(&patient)
	if patient.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//TODO: check only if title is duplicate
	// loop, title matches with course.coursename, JSON

	// generate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	patient.PatientId = strconv.Itoa(rand.Intn(100))
	patients = append(patients, patient)
	json.NewEncoder(w).Encode(patient)
	return

}

func updateOnePatient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one patient")
	w.Header().Set("Content-Type", "applicatioan/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID

	for index, patient := range patients {
		if patient.PatientId == params["id"] {
			patients = append(patients[:index], patients[index+1:]...)
			var patient Patient
			_ = json.NewDecoder(r.Body).Decode(&patient)
			patient.PatientId = params["id"]
			patients = append(patients, patient)
			json.NewEncoder(w).Encode(patient)
			return
		}
	}
	//TODO: send a response when id is not found
}

func deleteOnePatient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one patient")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)

	//loop, id, remove (index, index+1)

	for index, patient := range patients {
		if patient.PatientId == params["id"] {
			patients = append(patients[:index], patients[index+1:]...)
			// TODO: send a confirm or deny response
			break
		}
	}
}
